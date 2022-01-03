module.exports = {
  root: true,
  env: {
    node: true,
  },
  extends: [
    'plugin:vue/vue3-essential',
    'eslint:recommended',
    '@vue/typescript/recommended',
    '@vue/prettier',
    '@vue/prettier/@typescript-eslint',
  ],
  parserOptions: {
    ecmaVersion: 2020,
  },
  rules: {
    'no-console': 'off',
    'no-debugger': 'off',
    'no-empty': 'off',
    '@typescript-eslint/no-empty-function': 'off',
    'vue/no-use-v-if-with-v-for': 'off',
    '@typescript-eslint/no-unused-vars': 'off',
    'vue/no-unused-components': 'off',
    '@typescript-eslint/ban-ts-comment': 'off',
    'vue/no-mutating-props': 'off',
    '@typescript-eslint/no-var-requires': 'off',
  },
  overrides: [
    {
      files: ['**/__tests__/*.{j,t}s?(x)', '**/tests/unit/**/*.spec.{j,t}s?(x)'],
      env: {
        mocha: true,
      },
    },
  ],
};
