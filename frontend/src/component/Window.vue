<template>
  <div
    class="window"
    :style="{
      left: modelValue.x + 'px',
      top: modelValue.y + 'px',
      width: modelValue.width + 'px',
      height: modelValue.height + 'px',
      zIndex: $root.topPid == modelValue.pid ? 10 : 1,
    }"
  >
    <div draggable="false" ref="header" class="header" @mousedown="down">
      <div>{{ modelValue.title }}</div>
      <div style="margin-left: auto">
        <img
          @click="$emit('close', modelValue.pid)"
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
    <div draggable="false" @mousedown="downDrag($event, 'T')" class="t"></div>
    <div draggable="false" @mousedown="downDrag($event, 'B')" class="b"></div>
    <div draggable="false" @mousedown="downDrag($event, 'L')" class="l"></div>
    <div draggable="false" @mousedown="downDrag($event, 'R')" class="r"></div>
    <div
      draggable="false"
      @mousedown="downDrag($event, 'R'), downDrag($event, 'B')"
      class="rb"
    ></div>
    <div
      draggable="false"
      @mousedown="downDrag($event, 'L'), downDrag($event, 'B')"
      class="lb"
    ></div>
    <div
      draggable="false"
      @mousedown="downDrag($event, 'R'), downDrag($event, 'T')"
      class="tr"
    ></div>
    <div
      draggable="false"
      @mousedown="downDrag($event, 'L'), downDrag($event, 'T')"
      class="tl"
    ></div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { RestApi } from '../util/RestApi';

export default defineComponent({
  props: {
    modelValue: Object,
  },
  components: {},
  async mounted() {
    this.documentMove = (e: MouseEvent) => {
      if (!this.modelValue) {
        return;
      }

      if (this.isDrag) {
        this.$emit('update:modelValue', {
          ...this.modelValue,
          x: this.modelValue.x + (e.pageX - this.capture.x),
          y: this.modelValue.y + (e.pageY - this.capture.y),
        });
      }

      if (this.isDragB) {
        this.$emit('update:modelValue', {
          ...this.modelValue,
          height: this.modelValue.height + (e.pageY - this.capture.y),
        });
      }

      if (this.isDragT) {
        this.$emit('update:modelValue', {
          ...this.modelValue,
          y: this.modelValue.y + (e.pageY - this.capture.y),
          height: this.modelValue.height - (e.pageY - this.capture.y),
        });
      }

      if (this.isDragR) {
        this.$emit('update:modelValue', {
          ...this.modelValue,
          width: this.modelValue.width + (e.pageX - this.capture.x),
        });
      }

      if (this.isDragL) {
        this.$emit('update:modelValue', {
          ...this.modelValue,
          x: this.modelValue.x + (e.pageX - this.capture.x),
          width: this.modelValue.width - (e.pageX - this.capture.x),
        });
      }

      if (this.isDragR && this.isDragB) {
        this.$emit('update:modelValue', {
          ...this.modelValue,
          width: this.modelValue.width + (e.pageX - this.capture.x),
          height: this.modelValue.height + (e.pageY - this.capture.y),
        });
      }

      if (this.isDragL && this.isDragB) {
        this.$emit('update:modelValue', {
          ...this.modelValue,
          x: this.modelValue.x + (e.pageX - this.capture.x),
          width: this.modelValue.width - (e.pageX - this.capture.x),
          height: this.modelValue.height + (e.pageY - this.capture.y),
        });
      }

      if (this.isDragR && this.isDragT) {
        this.$emit('update:modelValue', {
          ...this.modelValue,
          width: this.modelValue.width + (e.pageX - this.capture.x),
          y: this.modelValue.y + (e.pageY - this.capture.y),
          height: this.modelValue.height - (e.pageY - this.capture.y),
        });
      }

      if (this.isDragL && this.isDragT) {
        this.$emit('update:modelValue', {
          ...this.modelValue,
          y: this.modelValue.y + (e.pageY - this.capture.y),
          height: this.modelValue.height - (e.pageY - this.capture.y),
          x: this.modelValue.x + (e.pageX - this.capture.x),
          width: this.modelValue.width - (e.pageX - this.capture.x),
        });
      }

      if (!(this.isDrag || this.isDragB || this.isDragT || this.isDragL || this.isDragR)) {
        return;
      }

      (this.$root as any).isAnyDrag = true;

      this.capture = {
        x: e.pageX,
        y: e.pageY,
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

    document.addEventListener('mousemove', this.documentMove);
    document.addEventListener('mouseup', this.documentUp);
  },
  beforeUnmount() {
    document.removeEventListener('mousemove', this.documentMove);
    document.removeEventListener('mouseup', this.documentUp);
  },
  methods: {
    down(e: any) {
      this.isDrag = true;
      this.capture = {
        x: e.pageX,
        y: e.pageY,
      };
      (this.$root as any).topPid = this.modelValue?.pid;
    },
    downDrag(e: any, dir: string) {
      this.capture = {
        x: e.pageX,
        y: e.pageY,
      };
      // @ts-ignore
      this[`isDrag${dir}`] = true;
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
  .rb,
  .lb,
  .tr,
  .tl {
    position: absolute;
    // background: #fe0000;
  }

  .t {
    height: 8px;
    width: 100%;
    top: -8px;
    cursor: s-resize;
  }

  .b {
    height: 8px;
    width: 100%;
    bottom: -8px;
    cursor: s-resize;
  }

  .l {
    height: 100%;
    width: 8px;
    left: -8px;
    top: 0;
    cursor: e-resize;
  }

  .r {
    height: 100%;
    width: 8px;
    right: -8px;
    top: 0;
    cursor: e-resize;
  }

  .rb {
    height: 8px;
    width: 8px;
    right: -8px;
    bottom: -8px;
    cursor: nw-resize;
  }

  .lb {
    height: 8px;
    width: 8px;
    left: -8px;
    bottom: -8px;
    cursor: ne-resize;
  }

  .tr {
    height: 8px;
    width: 8px;
    right: -8px;
    top: -8px;
    cursor: ne-resize;
  }

  .tl {
    height: 8px;
    width: 8px;
    left: -8px;
    top: -8px;
    cursor: nw-resize;
  }
}
</style>
