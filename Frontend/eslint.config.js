import globals from "globals";
import pluginJs from "@eslint/js";
import pluginVue from "eslint-plugin-vue";
import stylisticJs from "@stylistic/eslint-plugin-js";


export default [
  {files: ["**/*.{js,mjs,cjs,vue}"]},
  {languageOptions: { globals: globals.browser }},
  pluginJs.configs.recommended,
  ...pluginVue.configs["flat/essential"],
  {
    plugins: {
      '@stylistic/js': stylisticJs
    },
    rules: {
      '@stylistic/js/no-extra-semi': ['error'],
      '@stylistic/js/semi': ['error'],
      '@stylistic/js/semi-style': ['error']
    }
  }
];