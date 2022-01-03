<template>
  <div>
    <div class="bottom">
      <!-- Open -->
      <ui-button
        class="controlPanelButton"
        @click.stop="isShowPanel = !isShowPanel"
        icon="control_panel"
        :icon-size="32"
      />

      <!-- Multiple -->
      <div class="multiple">
        <div
          class="item clickable"
          :class="x - 1 === desktopId ? 'selected' : ''"
          v-for="x in 4"
          :key="x"
          @click="$emit('changeDesktop', x - 1)"
        ></div>
      </div>

      <!-- Items -->
      <div
        @click="toggle(x)"
        class="clickable icon"
        :class="x.pid === $root.topPid ? 'selected' : ''"
        v-for="x in windowList"
        v-show="x.desktopId === desktopId"
        :key="x.pid"
      >
        <img
          :src="`${$root.API_URL}/application/icon?path=${x.path}`"
          :alt="x.name"
          draggable="false"
        />
        <div>{{ x.name.replace(/_/g, ' ') }}</div>
      </div>

      <!-- Time -->
      <div class="time">
        <div>
          <b>{{ date }}</b>
        </div>
        <div>{{ time }}</div>
      </div>

      <!-- Info -->
      <ui-button
        class="infoButton"
        @click.stop="isShowHelp = !isShowHelp"
        icon="info"
        :icon-size="24"
        icon-color="#6b93ed"
      />
    </div>

    <!-- Panel -->
    <div v-if="isShowPanel" class="panel">
      <div
        @click="
          $emit('run', x.author + '/' + x.name + '@' + x.version);
          isShowPanel = false;
        "
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
          {{ x.name.replace(/_/g, ' ') }}
          <div class="version">{{ x.version }}</div>
        </div>
      </div>
    </div>

    <!-- Help -->
    <div v-if="isShowHelp" class="help">
      <div class="window">
        <div class="content">
          <h1>Desktop Application</h1>
          <p>You can open and use GAM application.</p>
          <h2>API</h2>
          <p><b>Signal Api</b> - You can send some data to specific application through Desktop.</p>
          <p>
            Send example JSON <code>appId: "maldan/fileman", data: "Data"</code> to
            <i>http://{{ currentHost }}/api/main/applicationData</i>
          </p>
          <p>
            Then desktop will save this data to buffer and then try to send to application process.
            If it fails to send it will try to send each time you run application. Desktop will send
            to application the path to JSON file with data. Application must delete this file if
            success.
          </p>
        </div>
        <ui-button @click="isShowHelp = !isShowHelp" text="Close" style="padding: 5px" />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { RestApi } from '@/util/RestApi';
import Moment from 'moment';

export default defineComponent({
  props: {
    windowList: Array,
    applicationList: Array,
    desktopId: Number,
  },
  components: {},
  async mounted() {
    document.addEventListener('click', () => {
      this.isShowPanel = false;
    });

    this.date = Moment().format('DD MMM YY');
    this.time = Moment().format('HH:mm:ss');
    this.timer = setInterval(() => {
      this.date = Moment().format('DD MMM YY');
      this.time = Moment().format('HH:mm:ss');
    }, 1000);
  },
  beforeUnmount() {
    clearInterval(this.timer);
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
      }
    },
  },
  data: () => {
    return {
      isShowPanel: false,
      isShowHelp: false,
      date: '',
      time: '',
      timer: null as any,
      currentHost: window.location.host,
    };
  },
});
</script>

<style lang="scss" scoped>
.panel {
  position: absolute;
  left: 0;
  bottom: 40px;
  background: rgba(0, 0, 0, 0.8);
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

.help {
  position: absolute;
  left: 0;
  top: 0;
  background: rgba(0, 0, 0, 0.4);
  width: 100%;
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 10000;

  .window {
    background: #2b2b2b;
    padding: 10px;
    max-width: 50%;

    .content {
      color: #a1a1a1;
      margin-bottom: 10px;
    }
  }
}

.bottom {
  position: fixed;
  bottom: 0;
  width: 100%;
  height: 40px;
  background: rgba(0, 0, 0, 0.8);
  backdrop-filter: blur(8px);
  display: flex;
  align-items: center;
  z-index: 20;
  user-select: none;

  .controlPanelButton {
    flex: 0;
    background: none;
    padding: 0 5px 0 15px;
  }

  .infoButton {
    flex: 0;
    background: none;
    padding: 0 15px 0 0;
  }

  .control_panel {
    padding-right: 20px;
    padding-left: 20px;
  }

  .multiple {
    margin-right: 20px;
    margin-left: 10px;
    display: grid;
    grid-template-columns: 1fr 1fr;
    gap: 3px;

    .item {
      width: 18px;
      height: 13px;
      background: rgba(255, 255, 255, 0.25);
      border-radius: 2px;

      &.selected {
        background: rgba(255, 255, 255, 0.5);
      }
    }
  }

  .time {
    margin-left: auto;
    padding-right: 15px;
    color: #bbbbbb;
    font-size: 13px;
    text-align: center;
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
