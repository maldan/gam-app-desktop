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
  >
    <div draggable="false" ref="header" class="header">
      <div>{{ modelValue.title }}</div>

      <div style="margin-left: auto; display: flex">
        <div
          @mousedown.stop=""
          @click.stop="minimize()"
          class="clickable"
          style="margin-right: 10px"
        >
          _
        </div>
        <div
          @mousedown.stop=""
          @click.stop="maximize()"
          class="clickable"
          style="margin-right: 10px"
        >
          []
        </div>
        <img
          @mousedown.stop=""
          @click.stop="$emit('close', modelValue.pid)"
          class="clickable"
          src="../asset/close.svg"
          alt=""
          draggable="false"
        />
      </div>
    </div>
    <div draggable="false" class="body">
      <div v-if="$root.isAnyDrag" class="overflow"></div>
      <iframe :src="modelValue.url" frameborder="0"></iframe>
    </div>

    <div v-for="x in resize" :key="x" :ref="`resize_${x}`" draggable="false" :class="x"></div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { RestApi } from '../util/RestApi';
import { Helper } from '../util/Helper';

export default defineComponent({
  props: {
    modelValue: Object,
  },
  components: {},
  async mounted() {
    this.documentMove = (e: any) => {
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
    };

    this.documentUp = (e: MouseEvent) => {
      if (this.isDrag || this.isDragB || this.isDragT || this.isDragL || this.isDragR) {
        RestApi.process.setWindow(this.modelValue as any);
      }

      this.isDrag = false;
      this.isDragT = false;
      this.isDragB = false;
      this.isDragL = false;
      this.isDragR = false;
      (this.$root as any).isAnyDrag = false;
    };

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
  },
  beforeUnmount() {
    Helper.removeEventListener(document, 'mousemove touchmove', this.documentMove);
    Helper.removeEventListener(document, 'mouseup touchend', this.documentUp);
  },
  methods: {
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

        this.$emit('update:modelValue', {
          ...this.modelValue,
          //isMinimized: false,
          isMaximized: isMaximized,
        });
        await RestApi.process.setWindow({
          ...(this.modelValue as any),
          //isMinimized: false,
          isMaximized: isMaximized,
        });
      }
    },
    async minimize() {
      if (this.modelValue) {
        this.$emit('update:modelValue', {
          ...this.modelValue,
          isMinimized: true,
          //isMaximized: false,
        });
        await RestApi.process.setWindow({
          ...(this.modelValue as any),
          isMinimized: true,
          //isMaximized: false,
        });
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

      capture: {
        x: 0,
        y: 0,
      },

      documentMove: null as any,
      documentUp: null as any,

      resize: ['l', 'r', 't', 'b', 'tr', 'tl', 'br', 'bl'],
    };
  },
});
</script>

<style lang="scss" scoped>
.window {
  position: absolute;
  background: #2b2b2b46;
  backdrop-filter: blur(8px);
  border-radius: 8px;

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
