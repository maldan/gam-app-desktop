<template>
  <router-view />
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { Helper } from './util/Helper';
import Moment from 'moment';

export default defineComponent({
  components: {},
  async mounted() {},
  methods: {
    format(t: string) {
      // Code
      const codeList = [] as string[];
      t = t.replace(/```(.*?)```/gms, (r1, r2) => {
        codeList.push(r2);
        return `__code${codeList.length - 1}__`;
      });

      const lines = t.split('\n');
      const liList = [];

      for (let i = 0; i < lines.length; i++) {
        // Replace links
        lines[i] = lines[i].replace(/(https?:\/\/\S*)/g, (r1, r2) => {
          if (r2.match(/\.(png|jpeg|jpg|bmp|gif)$/)) {
            return `<img src="${r2}" style="max-width: 100%;" />`;
          }
          let youtube = r2.match(/https?:\/\/(?:www\.)?youtube\.com\/watch\?v=(.*)/);
          if (!youtube) {
            youtube = r2.match(/https?:\/\/youtu\.be\/(.*)/);
          }
          if (youtube) {
            return `<a target="_blank" href="https://www.youtube.com/watch?v=${youtube[1]}"><img src="https://img.youtube.com/vi/${youtube[1]}/0.jpg"/></a>`;
            // return `<iframe width="560" height="315" src="https://www.youtube.com/embed/${youtube[1]}" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>`;
          }
          return `<a target="_blank" href="${r2}">${r2}</a> `;
        });

        // Replace other
        lines[i] = lines[i].replace(/(#+) (.*?)$/g, (r1, r2, r3) => {
          const number = r2.length;
          return `<h${number}>${r3}</h${number}>`;
        });

        // Replace other
        lines[i] = lines[i].replace(/\*\*(.*?)\*\*/g, (r1, r2, r3) => {
          return `<b>${r2}</b>`;
        });

        // Replace other
        lines[i] = lines[i].replace(/^\* (.*?)$/g, (r1, r2, r3) => {
          liList.push(r2);
          return `<li>${r2}</li>`;
        });
      }

      let final = lines.map((x) => `<div class="line">${x || '&nbsp;'}</div>`).join('');
      final = final.replace(/__code(\d+)__/g, (r1, r2) => {
        return `<pre><code>${codeList[Number(r2)]}</code></pre>`;
      });
      return final;
    },
  },
  data: () => {
    return {
      isAnyDrag: false,
      topPid: 0,
      convertName: Helper.convertName,
      // @ts-ignore
      API_URL: process.env.VUE_APP_API_URL || `${window.location.origin}/api`,
      moment: Moment,
    };
  },
});
</script>

<style lang="scss" scoped>
.main {
}
</style>
