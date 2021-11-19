<template>
  <!-- <img alt="Vue logo" src="./assets/logo.png"> -->
  <!-- <HelloWorld msg="AAAAA"/> -->
  <Feed />
  <News />
</template>

<script>
// import HelloWorld from './components/HelloWorld.vue';
// import News from './components/News.vue';
import News from './components/News.vue';
import Feed from './components/Feed.vue';

export default {
  name: 'App',
  components: {
    // HelloWorld,
    News,
    Feed,
  }
}
</script>

<style>
/* #app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
} */
</style>
