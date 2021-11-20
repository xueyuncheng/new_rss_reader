<template>
  <div class="feed" v-for="feed in feeds" :key="feed.id">
    <input
      type="checkbox"
      v-model="checked_ids"
      :key="feed.id"
      :value="feed.id"
    />
    <label for="feed.id">{{ feed.name }}</label>
    <button @click="deleteFeed(feed.id)" :key="feed.id">删除</button>
  </div>

  <form @submit.prevent="addFeed(name)">
    <input type="text" v-model="name" />
    <button type="submit">Add</button>
  </form>

  <button @click="listNews(checked_ids)">刷新</button>

  <div class="news">
    <ol>
      <li v-for="item in newses" :key="item.id">
        <a :href="item.link">{{ item.title }}</a>
        [{{ item.publish_time }}]
      </li>
    </ol>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "Feed",

  data() {
    return {
      feeds: [],
      name: "",
      checked_ids: [],
      newses: [],
    };
  },

  created() {
    this.listFeed();
    this.listNews();
  },

  methods: {
    listFeed() {
      console.log(this.checked_ids);
      axios
        .get("http://localhost:10001/api/feeds")
        .then((response) => {
          this.feeds = response.data.data;
        })
        .catch((error) => {
          console.log(error);
        });
    },

    addFeed(name) {
      axios
        .post("http://localhost:10001/api/feeds", {
          name: name,
        })
        .then(() => {
          this.listFeed();
          this.name = "";
        })
        .catch((error) => {
          console.log(error);
        });
    },

    deleteFeed(feed_id) {
      axios
        .delete(`http://localhost:10001/api/feeds/${feed_id}`)
        .then(() => {
          this.listFeed();
        })
        .catch((error) => {
          console.log(error);
        });
    },

    listNews(checked_ids) {
      console.log(checked_ids);
      axios
        .get(`http://localhost:10001/api/news`, {
          params: {
            feed_ids: this.checked_ids,
          },
        })
        .then((response) => {
          this.newses = response.data.data;
        })
        .catch((error) => {
          console.log(error);
        });
    },
  },
};
</script>