<template>
  <div
    class="window"
    :style="{
      left: modelValue.isMaximized ? `0px` : modelValue.x + 'px',
      top: modelValue.isMaximized ? `0px` : modelValue.y + 'px',
      width: modelValue.isMaximized ? `100%` : modelValue.width + 'px',
      height: modelValue.isMaximized ? `calc(100% - 40px)` : modelValue.height + 'px',
      zIndex: $root.topPid == modelValue.pid ? 10 : 1,
    }"
    :class="modelValue.isMaximized ? 'maximized' : ''"
    v-if="!modelValue.isMinimized"
  >
    <div draggable="false" ref="header" class="header">
      <div class="title">
        <img
          :src="`${$root.API_URL}/application/icon?path=${modelValue.path}`"
          :alt="modelValue.name"
          draggable="false"
          style="max-height: 16px"
        />
        <div>{{ $root.convertName(modelValue.appId) }}</div>
      </div>
      <div style="margin-left: auto; display: flex">
        <img
          @mousedown.stop=""
          @click.stop="openSettings()"
          class="clickable"
          src="../asset/settings.svg"
          alt=""
          draggable="false"
        />
        <img
          @mousedown.stop=""
          @click.stop="minimize()"
          class="clickable"
          src="../asset/minimize.svg"
          alt=""
          draggable="false"
          style="margin-left: 20px"
        />
        <img
          @mousedown.stop=""
          @click.stop="maximize()"
          class="clickable"
          src="../asset/maximize.svg"
          alt=""
          draggable="false"
          style="margin-left: 20px"
        />
        <img
          @mousedown.stop=""
          @click.stop="$emit('close', modelValue.pid)"
          class="clickable"
          src="../asset/close.svg"
          alt=""
          draggable="false"
          style="margin-left: 20px"
        />
      </div>
    </div>
    <div draggable="false" class="body">
      <div v-if="$root.isAnyDrag" class="overflow"></div>
      <iframe v-show="!isOpenSettings" :src="modelValue.url" frameborder="0"></iframe>
      <div class="settings" v-if="isOpenSettings">
        <JsonEditor :input="config" />
        <button class="button clickable" @click="saveConfig()">Save config</button>
      </div>
    </div>

    <div v-for="x in resize" :key="x" :ref="`resize_${x}`" draggable="false" :class="x"></div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { RestApi } from '../util/RestApi';
import { Helper } from '../util/Helper';
import Axios from 'axios';
import JsonEditor from './JsonEditor.vue';

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
          x: this.modelValue.x + (pageX - this.capture.x),
          y: this.modelValue.y + (pageY - this.capture.y),
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
          width: this.modelValue.width + (pageX - this.capture.x),
        });
      }

      if (this.isDragL) {
        this.$emit('update:modelValue', {
          ...this.modelValue,
          x: this.modelValue.x + (pageX - this.capture.x),
          width: this.modelValue.width - (pageX - this.capture.x),
        });
      }

      if (this.isDragR && this.isDragB) {
        this.$emit('update:modelValue', {
          ...this.modelValue,
          width: this.modelValue.width + (pageX - this.capture.x),
          height: this.modelValue.height + (pageY - this.capture.y),
        });
      }

      if (this.isDragL && this.isDragB) {
        this.$emit('update:modelValue', {
          ...this.modelValue,
          x: this.modelValue.x + (pageX - this.capture.x),
          width: this.modelValue.width - (pageX - this.capture.x),
          height: this.modelValue.height + (pageY - this.capture.y),
        });
      }

      if (this.isDragR && this.isDragT) {
        this.$emit('update:modelValue', {
          ...this.modelValue,
          width: this.modelValue.width + (pageX - this.capture.x),
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
          width: this.modelValue.width - (pageX - this.capture.x),
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
      try {
        this.config = (await Axios.get(this.modelValue?.url + 'system/config')).data.response;
      } catch {
        this.config = {};
      }
    },
    async saveConfig() {
      try {
        await Axios.post(this.modelValue?.url + 'system/config', this.config);
      } catch (e) {
        alert(e);
      }
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

      capture: {
        x: 0,
        y: 0,
      },

      //documentMove: null as any,
      //documentUp: null as any,

      resize: ['l', 'r', 't', 'b', 'tr', 'tl', 'br', 'bl'],
      config: {},
      dataOutput: {},
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
    overflow: hidden;

    .title {
      text-transform: capitalize;
      display: flex;
      align-items: center;

      > div {
        margin-left: 10px;
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
