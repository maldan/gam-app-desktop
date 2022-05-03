<template>
  <div
    class="window"
    :style="windowStyle()"
    :class="modelValue.dock === 'full' ? 'maximized' : ''"
    v-if="!modelValue.isMinimized"
  >
    <div draggable="false" ref="header" class="header">
      <!-- Title -->
      <div class="title">
        <img
          :src="`${$root.API_URL}/application/icon?path=${modelValue.path}`"
          :alt="modelValue.name"
          draggable="false"
          style="max-height: 16px"
        />
        <div>{{ modelValue.name.replace(/_/g, ' ') }}</div>
      </div>

      <div style="margin-left: auto; display: flex">
        <ui-button
          @mousedown.stop=""
          @click.stop="openSettings()"
          class="button"
          icon="gear"
          icon-color="#7e7e7e"
        />
        <ui-button
          @mousedown.stop=""
          @click.stop="openInfo()"
          class="button"
          icon="info"
          icon-color="#4672b5"
        />
        <ui-button @mousedown.stop="" @click.stop="minimize()" class="button" icon="minimize" />
        <ui-button
          @mousedown.stop=""
          @click.stop="isShowDockOption = !isShowDockOption"
          class="button"
          icon="maximize"
          icon-color="#2f8b0a"
        />
        <ui-button
          @mousedown.stop=""
          @click.stop="$emit('close', modelValue.pid)"
          class="button"
          icon="close"
          icon-color="#b12929"
        />

        <div class="dock_option" v-if="isShowDockOption">
          <button
            v-for="x in dockOptionList"
            :key="x.value"
            @click.stop="setDock(x.value)"
            class="clickable"
          >
            <ui-icon :name="x.icon" :color="x.value === modelValue.dock ? '#00ff00' : '#999999'" />
          </button>
        </div>
      </div>
    </div>
    <div draggable="false" class="body">
      <div v-if="$root.isAnyDrag" class="overflow"></div>
      <iframe
        v-show="!isOpenSettings && !isOpenInfo"
        :src="modelValue.url"
        frameborder="0"
      ></iframe>
      <div class="settings" v-if="isOpenSettings">
        <JsonEditor :input="config" />
        <button class="button clickable" @click="saveConfig()">Save config</button>
      </div>

      <!-- Info -->
      <div class="settings" v-if="isOpenInfo">
        <div>{{ modelValue.url }}</div>
        <div v-html="$root.format(info)"></div>
      </div>
    </div>

    <div v-for="x in resize" :key="x" :ref="`resize_${x}`" draggable="false" :class="x"></div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { RestApi } from '@/util/RestApi';
import { Helper } from '@/util/Helper';
import JsonEditor from './JsonEditor.vue';
import Axios from 'axios';
declare const require: (path: string) => void;

export default defineComponent({
  props: {
    modelValue: Object,
  },
  components: { JsonEditor },
  async mounted() {
    this.init();
  },
  beforeUnmount() {
    this.uninit();
  },
  watch: {
    'modelValue.isMinimized'() {
      this.init();
    },
  },
  methods: {
    documentUp(e: MouseEvent) {
      if (this.isDrag || this.isDragB || this.isDragT || this.isDragL || this.isDragR) {
        RestApi.process.setWindow(this.modelValue as any);
      }

      this.isDrag = false;
      this.isDragT = false;
      this.isDragB = false;
      this.isDragL = false;
      this.isDragR = false;
      (this.$root as any).isAnyDrag = false;
    },
    documentMove(e: any) {
      if (!this.modelValue) {
        return;
      }

      if (this.modelValue.isMaximized) {
        return;
      }

      const pageX = e.changedTouches ? e.changedTouches[0].pageX : e.pageX;
      const pageY = e.changedTouches ? e.changedTouches[0].pageY : e.pageY;

      if (this.isDrag) {
        this.$emit('update:modelValue', {
          ...this.modelValue,
          x: Math.max(this.modelValue.x + (pageX - this.capture.x), 0),
          y: Math.max(this.modelValue.y + (pageY - this.capture.y), 0),
        });
      }

      if (this.isDragB) {
        this.$emit('update:modelValue', {
          ...this.modelValue,
          height: this.modelValue.height + (pageY - this.capture.y),
        });
      }

      if (this.isDragT) {
        this.$emit('update:modelValue', {
          ...this.modelValue,
          y: this.modelValue.y + (pageY - this.capture.y),
          height: this.modelValue.height - (pageY - this.capture.y),
        });
      }

      if (this.isDragR) {
        this.$emit('update:modelValue', {
          ...this.modelValue,
          width: Math.max(this.modelValue.width + (pageX - this.capture.x), 320),
        });
      }

      if (this.isDragL) {
        this.$emit('update:modelValue', {
          ...this.modelValue,
          x: this.modelValue.x + (pageX - this.capture.x),
          width: Math.max(this.modelValue.width - (pageX - this.capture.x), 320),
        });
      }

      if (this.isDragR && this.isDragB) {
        this.$emit('update:modelValue', {
          ...this.modelValue,
          width: Math.max(this.modelValue.width + (pageX - this.capture.x), 320),
          height: this.modelValue.height + (pageY - this.capture.y),
        });
      }

      if (this.isDragL && this.isDragB) {
        this.$emit('update:modelValue', {
          ...this.modelValue,
          x: this.modelValue.x + (pageX - this.capture.x),
          width: Math.max(this.modelValue.width - (pageX - this.capture.x), 320),
          height: this.modelValue.height + (pageY - this.capture.y),
        });
      }

      if (this.isDragR && this.isDragT) {
        this.$emit('update:modelValue', {
          ...this.modelValue,
          width: Math.max(this.modelValue.width + (pageX - this.capture.x), 320),
          y: this.modelValue.y + (pageY - this.capture.y),
          height: this.modelValue.height - (pageY - this.capture.y),
        });
      }

      if (this.isDragL && this.isDragT) {
        this.$emit('update:modelValue', {
          ...this.modelValue,
          y: this.modelValue.y + (pageY - this.capture.y),
          height: this.modelValue.height - (pageY - this.capture.y),
          x: this.modelValue.x + (pageX - this.capture.x),
          width: Math.max(this.modelValue.width - (pageX - this.capture.x), 320),
        });
      }

      if (!(this.isDrag || this.isDragB || this.isDragT || this.isDragL || this.isDragR)) {
        return;
      }

      (this.$root as any).isAnyDrag = true;

      this.capture = {
        x: pageX,
        y: pageY,
      };
    },
    windowStyle() {
      const model = this.modelValue || {};
      // @ts-ignore
      const topPid = this.$root.topPid;
      const isDocked = model.dock !== '';

      let left = isDocked ? `0px` : model.x + 'px';
      let top = isDocked ? `0px` : model.y + 'px';
      let width = isDocked ? `100%` : model.width + 'px';
      let height = isDocked ? `calc(100% - 40px)` : model.height + 'px';

      if (model.dock === 'left') {
        width = '50%';
      }

      if (model.dock === 'right') {
        left = '50%';
        width = '50%';
      }

      if (this.isMobile()) {
        left = '0';
        top = '0';
        width = '100%';
        height = 'calc(100% - 40px)';
      }

      return {
        left,
        top,
        width,
        height,
        zIndex: topPid == model.pid ? 10 : 1,
      };
    },
    init() {
      if (this.modelValue?.isMinimized) {
        return;
      }
      this.uninit();

      this.$nextTick(() => {
        Helper.addEventListener(document, 'mousemove touchmove', this.documentMove);
        Helper.addEventListener(document, 'mouseup touchend', this.documentUp);
        Helper.addEventListener(this.$refs['header'] as any, 'mousedown touchstart', this.down);

        this.resize.forEach((x: string) => {
          Helper.addEventListener(this.$refs[`resize_${x}`], 'mousedown touchstart', (e: any) => {
            const y = x.split('').map((xx) => xx.toUpperCase());
            y.forEach((yyy) => {
              this.downDrag(e, yyy);
            });
          });
        });
      });
    },
    uninit() {
      Helper.removeEventListener(document, 'mousemove touchmove', this.documentMove);
      Helper.removeEventListener(document, 'mouseup touchend', this.documentUp);
    },
    down(e: any) {
      const pageX = e.changedTouches ? e.changedTouches[0].pageX : e.pageX;
      const pageY = e.changedTouches ? e.changedTouches[0].pageY : e.pageY;

      this.isDrag = true;
      this.capture = {
        x: pageX,
        y: pageY,
      };
      (this.$root as any).topPid = this.modelValue?.pid;
    },
    downDrag(e: any, dir: string) {
      const pageX = e.changedTouches ? e.changedTouches[0].pageX : e.pageX;
      const pageY = e.changedTouches ? e.changedTouches[0].pageY : e.pageY;

      this.capture = {
        x: pageX,
        y: pageY,
      };
      // @ts-ignore
      this[`isDrag${dir}`] = true;
    },
    async maximize() {
      if (this.modelValue) {
        const isMaximized = !this.modelValue.isMaximized;

        await RestApi.process.setWindow({
          ...(this.modelValue as any),
          isMaximized: isMaximized,
        });
        this.$emit('update:modelValue', {
          ...this.modelValue,
          isMaximized: isMaximized,
        });
      }
    },
    async setDock(dock: string) {
      this.isShowDockOption = false;

      if (this.modelValue) {
        await RestApi.process.setWindow({
          ...(this.modelValue as any),
          dock,
        });
        this.$emit('update:modelValue', {
          ...this.modelValue,
          dock,
        });
      }
    },
    async minimize() {
      if (this.modelValue) {
        await RestApi.process.setWindow({
          ...(this.modelValue as any),
          isMinimized: true,
        });
        this.$emit('update:modelValue', {
          ...this.modelValue,
          isMinimized: true,
        });
      }
    },
    async openSettings() {
      this.isOpenSettings = !this.isOpenSettings;
      this.isOpenInfo = false;

      try {
        this.config = (await RestApi.application.getConfig(this.modelValue?.appId + '')) || {};
      } catch {
        this.config = {};
      }
    },
    async openInfo() {
      this.isOpenInfo = !this.isOpenInfo;
      this.isOpenSettings = false;
      this.info = '';
      this.info = (await Axios.get(`${this.modelValue?.url}system/info`)).data.response;
    },
    async saveConfig() {
      try {
        await RestApi.application.saveConfig(this.modelValue?.appId + '', this.config);
      } catch (e) {
        alert(e);
      }
    },
    isMobile() {
      return window.outerWidth <= 576;
    },
  },
  data: () => {
    return {
      isDrag: false,
      isDragT: false,
      isDragB: false,
      isDragL: false,
      isDragR: false,
      isOpenSettings: false,
      isShowDockOption: false,
      isOpenInfo: false,

      capture: {
        x: 0,
        y: 0,
      },

      dockOptionList: [
        { icon: 'maximize', value: '' },
        { icon: 'left_screen', value: 'left' },
        { icon: 'full_screen', value: 'full' },
        { icon: 'right_screen', value: 'right' },
      ],

      //documentMove: null as any,
      //documentUp: null as any,

      resize: ['l', 'r', 't', 'b', 'tr', 'tl', 'br', 'bl'],
      config: {},
      info: '',
    };
  },
});
</script>

<style lang="scss" scoped>
.window {
  position: absolute;
  background: rgba(0, 0, 0, 0.8);
  backdrop-filter: blur(8px);
  border-radius: 4px;
  user-select: none;

  &.maximized {
    border-radius: 0;
  }

  .header {
    user-select: none;
    padding: 10px;
    color: #afafaf;
    cursor: move;
    display: flex;
    align-items: center;
    white-space: nowrap;
    text-overflow: ellipsis;
    // overflow: hidden;
    position: relative;
    box-sizing: border-box;

    .button {
      padding: 0;
      background: none;
      margin-left: 5px;
    }

    .title {
      text-transform: capitalize;
      display: flex;
      align-items: center;

      > div {
        margin-left: 10px;
      }
    }

    .dock_option {
      position: absolute;
      right: 0px;
      bottom: -33px;
      background: #1b1b1b;
      z-index: 1;
      padding: 5px;
      box-sizing: border-box;

      button {
        background: none;
        border: 0;
        outline: none;
        color: #999999;
      }

      .selected {
        img {
          filter: invert(8%) sepia(100%) saturate(6505%) hue-rotate(124deg) brightness(118%)
            contrast(123%);
        }
      }
    }
  }

  .body {
    width: 100%;
    height: calc(100% - 40px);
    position: absolute;

    .overflow {
      position: absolute;
      left: 0;
      top: 40px;
      width: 100%;
      height: 100%;
      background: transparent;
    }

    iframe {
      display: block;
      width: 100%;
      height: 100%;
    }

    &::-webkit-scrollbar {
      width: 10px;
      background: rgba(0, 0, 0, 0.2);
    }

    &::-webkit-scrollbar-thumb {
      background: rgba(0, 0, 0, 0.5);
      border-radius: 50px;
    }

    .settings {
      display: flex;
      flex-direction: column;
      padding: 10px;
      overflow-y: auto;
      height: calc(100% - 25px);
      background: #2b2b2b;
      user-select: text;
      width: calc(100% - 20px);

      .textarea {
        height: 200px;
        resize: vertical;
      }

      button {
        margin-top: 10px;
      }
    }
  }

  .b,
  .t,
  .l,
  .r,
  .br,
  .bl,
  .tr,
  .tl {
    position: absolute;
    // background: #fe0000;
  }

  .t {
    height: 24px;
    width: 100%;
    top: -24px;
    cursor: s-resize;
  }

  .b {
    height: 24px;
    width: 100%;
    bottom: -24px;
    cursor: s-resize;
  }

  .l {
    height: 100%;
    width: 24px;
    left: -24px;
    top: 0;
    cursor: e-resize;
  }

  .r {
    height: 100%;
    width: 24px;
    right: -24px;
    top: 0;
    cursor: e-resize;
  }

  .br {
    height: 24px;
    width: 24px;
    right: -24px;
    bottom: -24px;
    cursor: nw-resize;
  }

  .bl {
    height: 24px;
    width: 24px;
    left: -24px;
    bottom: -24px;
    cursor: ne-resize;
  }

  .tr {
    height: 24px;
    width: 24px;
    right: -24px;
    top: -24px;
    cursor: ne-resize;
  }

  .tl {
    height: 24px;
    width: 24px;
    left: -24px;
    top: -24px;
    cursor: nw-resize;
  }
}
</style>
