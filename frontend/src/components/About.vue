<script setup lang="ts">
import {ref,watchEffect} from "vue"
import {useRoute} from "vue-router"

const loaded = ref(false)
const error = ref(false)
const product = ref(false)
const route = useRoute()
</script>

<template>
  <h1>About Me!</h1>
  <p>This is home page</p>
</template>