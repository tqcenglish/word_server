

// import index from "./page/index.vue"
import zh from "./i18n/zh.js"
import en from "./i18n/en.js"

const messages = {
  en,
  zh,
};

let locale = navigator.language;
if(locale != "zh"){
  locale = "en"
}

const i18n = new VueI18n({
  locale,
  messages, // set locale messages
});

new Vue({
  i18n,
  el: '#app',
  components: {
      index: httpVueLoader("./page/index.vue")
    },
  mounted(){
  }    
});