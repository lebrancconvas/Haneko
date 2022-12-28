<template>
  <div>
    <header>
      <div>
        <h1>Haneko</h1>
      </div>
      <div>
        <h2>Best Platform to start your blog</h2>
      </div>
    </header>
    <section id="action-section">
      <div>
        <NuxtLink to="/startblog">
          <Button id="start-blog-button" type="success">Start Your Blog</Button>
        </NuxtLink>
        <NuxtLink to="/bloglist">
          <Button id="blog-list-button" type="primary">Blog List</Button>
        </NuxtLink>
      </div>
    </section>
  </div>
</template>

<script>

</script>

<style scoped>
  * {
    padding: 0;
    margin: 0;
    text-align: center;
    vertical-align: center;
  }

  #action-section {
    margin: 30px;
  }

  #start-blog-button {
    margin: 10px;
    padding: 5px;
  }

  #blog-list-button {
    margin: 10px;
    padding: 5px;
  }
</style>
