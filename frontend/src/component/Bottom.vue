<template>
  <div class="bottom">
    <div
      @click="unminimize(x), top(x)"
      class="clickable icon"
      v-for="x in modelValue"
      :key="x.pid"
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
  async mounted() {},
  methods: {
    async unminimize(x: any) {
      x.isMinimized = false;
      await RestApi.process.setWindow({ ...(x as any), isMinimized: false });
    },
    top(x: any) {
      (this.$root as any).topPid = x.pid;
    },
  },
  data: () => {
    return {};
  },
});
</script>

<style lang="scss" scoped>
.bottom {
  position: fixed;
  bottom: 0;
  width: 100%;
  height: 40px;
  background: #2b2b2b40;
  backdrop-filter: blur(8px);
  display: flex;
  align-items: center;
  padding-left: 5px;
  z-index: 20;

  .icon {
    width: 32px;
    height: 32px;
    background: #fe0000;
    margin-right: 10px;
  }
}
</style>
