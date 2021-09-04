import Axios from 'axios';
import { config } from 'chai';

const API_URL = process.env.VUE_APP_API_URL || `${window.location.origin}/api`;

export const RestApi = {
  process: {
    async list() {
      return (await Axios.get(`${API_URL}/process/list`)).data.response;
    },
    async kill(pid: number) {
      return (
        await Axios.post(`${API_URL}/process/kill`, {
          pid,
        })
      ).data.response;
    },
    async run(url: string, host: string) {
      return (
        await Axios.post(`${API_URL}/process/run`, {
          url,
          host,
        })
      ).data.response;
    },
    async setWindow({
      pid,
      x,
      y,
      width,
      height,
      isMinimized,
      isMaximized,
    }: {
      pid: number;
      x: number;
      y: number;
      width: number;
      height: number;
      isMinimized: boolean;
      isMaximized: boolean;
    }) {
      return (
        await Axios.post(`${API_URL}/process/setWindow`, {
          pid,
          x,
          y,
          width,
          height,
          isMinimized,
          isMaximized,
        })
      ).data.response;
    },
  },

  application: {
    async list() {
      return (await Axios.get(`${API_URL}/application/list`)).data.response;
    },
    async getConfig(appId: string) {
      return (await Axios.get(`${API_URL}/application/config?appId=${appId}`)).data.response;
    },
    async saveConfig(appId: string, config: any = {}) {
      return (
        await Axios.post(`${API_URL}/application/config`, {
          appId,
          config,
        })
      ).data.response;
    },
  },
};
