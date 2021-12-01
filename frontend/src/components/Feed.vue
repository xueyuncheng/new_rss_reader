<template>
  <div class="feed_group">
    <form @submit.prevent="addFeed(name)" style="display: inline-block">
      <input type="text" v-model="name" placeholder="请输入RSS源" />
      <button type="submit">添加</button>
    </form>

    <button style="margin-left: 20px" @click="listNews()">刷新</button>

    <div class="feed" v-for="feed in feeds" :key="feed.id">
      <input
        type="radio"
        v-model="checked_id"
        :key="feed.id"
        :value="feed.id"
        :id="feed.id"
        @click="listNews(feed.id)"
      />
      <label :for="feed.id" style="margin-left: 5px">
        {{ feed.name }}
      </label>
      <button @click="deleteFeed(feed.id)" :key="feed.id">删除</button>
      <p style="display: inline-block; margin-left: 20px; color: red">
        {{ feed.error_msg }}
      </p>
    </div>
  </div>

  <div
    id="chart"
    style="width: 600px; height: 400px; position: absolute; top: 0; right: 0"
  ></div>

  <div class="news_group">
    <ol>
      <li class="news" v-for="item in newses" :key="item.id">
        <a :href="item.link" target="_blank">{{ item.title }}</a>
        [{{ item.publish_time }}]
      </li>
    </ol>
  </div>
</template>

<style>
button {
  min-width: 60px;
  min-height: 25px;
  margin-left: 20px;
}

.feed_group {
  margin-bottom: 30px;
}

.feed {
  height: 20px;
  margin-top: 10px;
}

.news {
  margin-top: 10px;
}
</style>

<script>
import axios from "axios";
import * as echarts from "echarts";

export default {
  name: "Feed",

  data() {
    return {
      feeds: [],
      name: "",
      checked_id: 0,
      newses: [],
    };
  },

  mounted() {
    this.initEchart();
  },

  created() {
    this.listFeed();
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

    initEchart() {
      axios.get(`/api/news/stat`).then((response) => {
        const data = response.data.data;
        const xAxis = [];
        const yAxis = [];

        data.items.forEach((element) => {
          xAxis.push(element.name);
          yAxis.push(element.value);
        });

        const myChart = echarts.init(document.getElementById("chart"));
        myChart.setOption({
          title: {
            text: response.data.name,
          },
          tooltip: {},
          xAxis: {
            data: xAxis,
          },
          yAxis: {},
          series: [
            {
              name: "新闻数量",
              type: "bar",
              data: yAxis,
            },
          ],
        });
      });
    },
  },
};
</script>

