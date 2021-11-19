<template>
    <div class="feed" v-for="feed in feeds" :key="feed.id">
            <input type="checkbox" v-model="feed.id" :key="feed.id"/>
            <label for="feed.id">{{ feed.name }}</label>
            <button @click="deleteFeed(feed.id)" :key="feed.id">
                删除 
            </button>
    </div>

    <form @submit.prevent="addFeed">
        <input type="text" v-model="name" />
        <button type="submit">Add</button>
    </form>

    <button @click="listFeed">刷新</button>
</template>

<script>
import axios from 'axios';

export default {
    name: 'Feed',
    data() {
        return {
            feeds: [],
            name: '',
        }
    },
    created() {
        this.listFeed();
    },
    methods: {
        listFeed() {
            axios.get('http://localhost:10001/api/feeds')
                .then(response => {
                    this.feeds = response.data.data;
                })
                .catch(error => {
                    console.log(error);
                });
        },

        addFeed() {
            axios.post('http://localhost:10001/api/feeds', {
                name: this.name,
            })
            .then(() => {
                this.listFeed();
                this.name = '';
            })
            .catch(error => {
                console.log(error);
            });
        },

        deleteFeed(feed_id) {
            axios.delete(`http://localhost:10001/api/feeds/${feed_id}`)
                .then(() => {
                    this.listFeed();
                })
                .catch(error => {
                    console.log(error);
                });
        },
    },
}
</script>