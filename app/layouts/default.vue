<template>
  <div>
    <nav>
      <NavBar />
    </nav>
    <br>
    <div>
      <nuxt />
    </div>
  </div>
</template>

<script>
  import NavBar from "@/components/NavBar";
  export default {
    components: {
      NavBar
    }
  }
</script>
