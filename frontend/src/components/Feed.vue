<template>
  <div class="feed_group">
    <form @submit.prevent="addFeed(name)" style="display: inline-block">
      <input type="text" v-model="name" placeholder="请输入RSS源" />
      <button type="submit">添加1</button>
    </form>

    <button style="margin-left: 20px" @click="listNews()">刷新</button>

    <div class="feed" v-for="feed in feeds" :key="feed.id">
      <input
        type="radio"
        v-model="checked_ids"
        :key="feed.id"
        :value="feed.id"
        :id="feed.id"
        @click="listNews(feed.id)"
      />
      <label :for="feed.id" style="margin-left: 5px">
        {{ feed.name }}
      </label>
      <button @click="deleteFeed(feed.id)" :key="feed.id">删除</button>
    </div>
  </div>

  <div class="news_group">
    <ol>
      <li class="news" v-for="item in newses" :key="item.id">
        <a :href="item.link" target="_blank">{{ item.title }}</a>
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
      axios
        .get("/api/feeds")
        .then((response) => {
          this.feeds = response.data.data;
        })
        .catch((error) => {
          console.log(error);
        });
    },

    addFeed(name) {
      axios
        .post("/api/feeds", {
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
        .delete(`api/feeds/${feed_id}`)
        .then(() => {
          this.listFeed();
        })
        .catch((error) => {
          console.log(error);
        });
    },

    listNews(feed_id) {
      axios
        .get(`/api/news`, {
          params: {
            feed_id: feed_id,
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

<style>
button {
  min-width: 60px;
  min-height: 25px;
  margin-left: 20px;
}

.feed {
  margin-top: 10px;
}

.news {
  margin-top: 5px;
}
</style>