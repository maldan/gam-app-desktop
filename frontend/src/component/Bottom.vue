<template>
  <div class="bottom"></div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  components: {},
  async mounted() {},
  methods: {},
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
}
</style>
