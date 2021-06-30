import Axios from 'axios';

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
    }: {
      pid: number;
      x: number;
      y: number;
      width: number;
      height: number;
    }) {
      return (
        await Axios.post(`${API_URL}/process/setWindow`, {
          pid,
          x,
          y,
          width,
          height,
        })
      ).data.response;
    },
  },

  application: {
    async list() {
      return (await Axios.get(`${API_URL}/application/list`)).data.response;
    },
  },
};
