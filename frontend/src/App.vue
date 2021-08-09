<template>
  <router-view />
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { Helper } from './util/Helper';

export default defineComponent({
  components: {},
  async mounted() {},
  methods: {},
  data: () => {
    return {
      isAnyDrag: false,
      topPid: 0,
      convertName: Helper.convertName,
      // @ts-ignore
      API_URL: process.env.VUE_APP_API_URL || `${window.location.origin}/api`,
    };
  },
});
</script>

<style lang="scss" scoped>
.main {
}
</style>
