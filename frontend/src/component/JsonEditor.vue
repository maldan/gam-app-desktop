<template>
  <div class="editor">
    <div class="line" v-for="(v, k) in input" :key="k">
      <div class="object">
        <div>{{ k }}</div>
        <input
          v-if="typeof v === 'string' || typeof v === 'number'"
          type="text"
          v-model="input[k]"
        />
        <div v-if="typeof v === 'object'">...</div>
        <div>{{ typeof v }}</div>
        <button @click="delete input[k]">x</button>
      </div>
      <JsonEditor v-if="typeof v === 'object'" :input="input[k]" />
    </div>

    <!-- Add new -->
    <div style="display: flex">
      <input type="text" placeholder="Key..." v-model="key" />
      <input type="text" placeholder="Value..." v-model="value" />
      <select v-model="type">
        <option value="string">String</option>
        <option value="number">Number</option>
        <option value="array">Array</option>
        <option value="object">Object</option>
      </select>
      <button @click="add(input, key, value, type)">+</button>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  props: {
    input: Object,
  },
  components: {},
  async mounted() {},
  methods: {
    add(to: any, key: string, value: any, type: string) {
      if (type === 'string') {
        to[key] = value;
      }
      if (type === 'number') {
        to[key] = Number(value);
      }
      if (type === 'array') {
        to[key] = [];
      }
      if (type === 'object') {
        to[key] = {};
      }

      this.key = '';
      this.value = '';
    },
  },
  data: () => {
    return {
      key: '',
      value: '',
      type: 'string',
    };
  },
});
</script>

<style lang="scss" scoped>
.editor {
  background: #1b1b1b;
  padding: 10px;
  border: 1px solid #6e6e6e;

  .line {
    .object {
      display: flex;
      box-sizing: border-box;
      color: #999999;
      margin-bottom: 10px;

      > div {
        box-sizing: border-box;
        flex: 1;
        padding: 5px;
      }
    }
  }

  input,
  select {
    background: #333333;
    color: #999999;
    padding: 5px 10px;
    outline: none;
    border: 0;
    flex: 1;
    font-size: 16px;
  }

  button {
    background: #3d3d3d;
    outline: none;
    border: none;
    border-radius: 2px;
    padding: 5px 10px;
    color: #999999;
    font-size: 16px;

    cursor: pointer;
    user-select: none;

    &:hover {
      opacity: 0.7;
    }

    &:active {
      position: relative;
      opacity: 0.6;
      top: 1px;
    }
  }
}
</style>
