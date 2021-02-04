import Vue from "vue";
import App from "./App.vue";
import router from "./router/router";
import '../../node_modules/materialize-css/dist/css/materialize.min.css';
import 'materialize-css/dist/js/materialize.min';

new Vue({
  el: "#app",
  router,
  render: h => h(App)
})