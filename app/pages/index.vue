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
          <Button type="success">Start Your Blog</Button>
        </NuxtLink>
        <NuxtLink to="/bloglist">
          <Button type="primary">Blog List</Button>
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
  }
</style>
