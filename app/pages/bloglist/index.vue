<template>
  <div>
    <header>
      <div>
        <h1>Blog List</h1>
      </div>
    </header>
    <section id="blog-list-section">
      <div v-for="(blog, index) in results" :key="index">
        <Card id="blog-list-card" style="width: 300px">
          <template #title>
            <h2>
              {{ blog.Title }}
            </h2>
          </template>
          <div>
            <h3>{{ blog.Description }}</h3>
            <Tag color="primary"> 
              <p>{{ blog.Category.toUpperCase() }}</p>
            </Tag>
          </div>
        </Card>
      </div>
    </section>
  </div>
</template>

<script>
  import axios from 'axios';
  export default {
    data() {
      return {
        results: []
      }
    },
    created() {
      axios.get("http://localhost:8000/api/v1/blogs")
        .then(res => {
          this.results = res.data;
          console.log(this.results);
        })
        .catch(err => {
          console.error(err);
          console.log(`[ERROR] Fetching API Error.`);
        })
    }
  }
</script>

<style scoped>
  * {
    text-align: center;
  }

  #blog-list-section {
    display: grid;
    grid-template-columns: repeat(3, 1fr);
  }

  #blog-list-card {
    margin: 10px;
    margin-left: auto;
    margin-right: auto;
    cursor: pointer;
  }
</style>
