<template>
  <div>
    <div class="bottom">
      <img
        @click.stop="isShowPanel = !isShowPanel"
        class="clickable control_panel"
        src="../asset/control_panel.svg"
        alt="Control Panel"
        draggable="false"
      />
      <div
        @click="toggle(x)"
        class="clickable icon"
        :class="x.pid === $root.topPid ? 'selected' : ''"
        v-for="x in windowList"
        :key="x.pid"
      >
        <img
          :src="`${$root.API_URL}/application/icon?path=${x.path}`"
          :alt="x.name"
          draggable="false"
        />
        <div>{{ $root.convertName(x.appId) }}</div>
      </div>
    </div>

    <!-- Panel -->
    <div v-if="isShowPanel" class="panel">
      <div
        @click="$emit('run', x.author + '/' + x.name + '@' + x.version), (isShowPanel = false)"
        class="app clickable"
        v-for="x in applicationList"
        :key="x.path"
      >
        <img
          :src="`${$root.API_URL}/application/icon?path=${x.path}`"
          :alt="x.name"
          draggable="false"
          style="width: 32px; max-height: 26px"
        />
        <div style="display: flex; flex: 1; align-items: center">
          {{ x.name }}
          <div class="version">{{ x.version }}</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { RestApi } from '../util/RestApi';

export default defineComponent({
  props: {
    windowList: Array,
    applicationList: Array,
  },
  components: {},
  async mounted() {
    document.addEventListener('click', () => {
      this.isShowPanel = false;
    });
  },
  methods: {
    async toggle(x: any) {
      if ((this.$root as any).topPid === x.pid) {
        x.isMinimized = !x.isMinimized;
        await RestApi.process.setWindow({ ...(x as any), isMinimized: x.isMinimized });
      } else {
        (this.$root as any).topPid = x.pid;
        if (x.isMinimized) {
          x.isMinimized = false;
          await RestApi.process.setWindow({ ...(x as any), isMinimized: false });
        }
        /*if (!x.isMinimized) {
          (this.$root as any).topPid = x.pid;
        }*/
        //
        //x.isMinimized = !x.isMinimized;
        //await RestApi.process.setWindow({ ...(x as any), isMinimized: x.isMinimized });
      }

      //x.isMinimized = !x.isMinimized;
      //await RestApi.process.setWindow({ ...(x as any), isMinimized: x.isMinimized });
      //(this.$root as any).topPid = x.pid;
    },

    /*convertName(name: string) {
      return name
        .split('-')
        .slice(1)
        .filter((x) => x !== 'app' && x !== 'gam')
        .join(' ');
    },*/
  },
  data: () => {
    return {
      isShowPanel: false,
    };
  },
});
</script>

<style lang="scss" scoped>
.panel {
  position: absolute;
  left: 0;
  bottom: 40px;
  background: rgba(0, 0, 0, 0.5);
  width: 320px;
  height: 480px;
  display: flex;
  flex-direction: column;
  overflow-y: auto;
  backdrop-filter: blur(8px);
  z-index: 30;

  .app {
    padding: 10px;
    color: #bbbbbb;
    display: flex;
    align-items: center;

    img {
      display: block;
      width: 32px;
      margin-right: 15px;
      max-height: 22px;
    }

    > div {
      overflow: hidden;
      text-overflow: ellipsis;
      white-space: nowrap;
      text-transform: capitalize;
    }

    .version {
      margin-left: auto;
      background: #1f1f1f73;
      font-size: 14px;
      padding: 5px 10px;
      border-radius: 4px;
      color: #9b9b9b;
      font-weight: bold;
    }
  }
}

.bottom {
  position: fixed;
  bottom: 0;
  width: 100%;
  height: 40px;
  background: rgba(0, 0, 0, 0.3);
  backdrop-filter: blur(8px);
  display: flex;
  align-items: center;
  // padding-left: 5px;
  z-index: 20;

  .control_panel {
    padding-right: 20px;
    padding-left: 20px;
  }

  .icon {
    //width: 32px;
    //height: 32px;
    //background: #fe0000;
    display: flex;
    margin-right: 5px;
    font-size: 14px;
    color: #c5c5c5;
    align-items: center;
    background: rgba(255, 255, 255, 0.1);
    height: 40px;
    box-sizing: border-box;
    padding: 0 10px;
    border-bottom: 2px solid rgba(255, 255, 255, 0.3);

    img {
      display: block;
      width: 32px;
      max-height: 22px;
    }

    > div {
      width: 100%;
      overflow: hidden;
      text-overflow: ellipsis;
      white-space: nowrap;
      margin-left: 10px;
      text-transform: capitalize;
    }

    &.selected {
      background: rgba(255, 255, 255, 0.3);
      border-bottom: 2px solid rgba(255, 255, 255, 0.8);
    }
  }
}
</style>
