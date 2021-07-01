<template>
  <div class="main">
    <Bottom @run="run($event)" :windowList="windowList" :applicationList="applicationList" />
    <Window
      @close="closeWindow"
      v-for="(x, i) in windowList"
      :key="x.pid"
      v-model="windowList[i]"
    />
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import Bottom from '../component/Bottom.vue';
import Window from '../component/Window.vue';
import { RestApi } from '../util/RestApi';

export default defineComponent({
  components: { Bottom, Window },
  async mounted() {
    this.applicationList = await RestApi.application.list();

    await this.refresh();
  },
  computed: {},

  methods: {
    async closeWindow(pid: number) {
      await RestApi.process.kill(pid);
      await this.refresh();
    },
    async refresh() {
      const list = await RestApi.process.list();
      this.windowList = list.map((p: any) => {
        return {
          pid: p.pid,
          x: 300,
          y: 200,
          width: 320,
          height: 240,
          title: p.args['app-id'],
          url: `http://${p.args.host}:${p.args.clientPort || p.args.port}/`,
          ...p.window,
        };
      });
    },
    async run(url: string) {
      await RestApi.process.run(url, window.location.hostname);
      await this.refresh();
    },
  },
  data: () => {
    return {
      applicationList: [] as any[],
      windowList: [
        // { id: 0, x: 100, y: 200, width: 1200, height: 820 }
      ] as any[],
    };
  },
});
</script>

<style lang="scss" scoped>
.main {
  position: absolute;
  left: 0;
  top: 0;
  background: url('https://cdn.pixabay.com/photo/2015/04/19/08/32/marguerite-729510_1280.jpg');
  width: 100%;
  height: 100%;
  background-size: cover;
  overflow: 'hidden';

  .app {
    user-select: none;
  }
}
</style>
