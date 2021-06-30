<template>
  <div class="main">
    <Bottom />
    <Window
      @close="closeWindow"
      v-for="(x, i) in windowList"
      :key="x.pid"
      v-model="windowList[i]"
    />
    <div class="app" @click="run(x.name)" v-for="x in applicationList" :key="x.path">
      {{ x.name }}
    </div>
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
    // this.run(`maldan/gam-app-caloryman`);
  },
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
          url: `http://${p.args.host}:${p.args.port}/`,
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
  position: fixed;
  left: 0;
  top: 0;
  background: url('../asset/bg.jpg');
  width: 100%;
  height: 100%;
  background-size: cover;

  .app {
    user-select: none;
  }
}
</style>
