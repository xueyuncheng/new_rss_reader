<template>
    <div class="news">
        <ol>
            <li v-for="item in newses" :key="item.id">
                <a :href="item.link">{{ item.title }}</a>
                [{{ item.publish_time}}] {{item.feed_name}}
            </li>
        </ol>
    </div>
</template>

<script>
import axios from 'axios';
export default {
    name: 'News',
    props: {
        // id: Number
    },
    data() {
        return {
            newses: [],
        }
    },
    created() {
        // axios.get(`https://jsonplaceholder.typicode.com/posts/${this.id}`)
        axios.get(`http://localhost:10001/news`)
            .then(response => {
                this.newses = response.data.data;
            }).catch(error => {
                console.log(error)
            })
    }
}
</script>