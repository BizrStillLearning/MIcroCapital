<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { Menu, X } from '@lucide/vue'
import { useAuthStore } from '../stores/auth.js'

const router = useRouter()
const authStore = useAuthStore()
const isOpen = ref(false)
const emit = defineEmits(['nav'])

const navItems = [
  { name: 'Campaigns', path: 'campaigns' },
  { name: 'How It Works', path: 'how-it-works' },
  { name: 'Services', path: 'services' },
  { name: 'Community', path: 'community' }
]

const navigate = (path) => {
  emit('nav', path)
}

const navigateMobile = (path) => {
  emit('nav', path)
  isOpen.value = false
}
</script>

<template>
  <header class="sticky top-0 z-50 bg-background/95 backdrop-blur-sm border-b border-border">
    <div class="max-w-6xl mx-auto px-5 h-16 flex items-center justify-between">

      <button
          @click="navigate('/')"
          class="flex items-center gap-2.5 group"
      >
        <div class="w-8 h-8 rounded-lg bg-primary flex items-center justify-center">
          <span class="text-primary-foreground font-bold text-sm" style="font-family: 'Fraunces', serif;">U</span>
        </div>
        <span class="font-semibold text-lg text-foreground" style="font-family: 'Fraunces', serif;">
          Umoja
        </span>
      </button>

      <nav class="hidden md:flex items-center gap-7 text-sm font-medium text-muted-foreground">
        <button
            v-for="item in navItems"
            :key="item.name"
            @click="navigate(item.path)"
            class="hover:text-foreground transition-colors"
        >
          {{ item.name }}
        </button>
      </nav>

      <div class="hidden md:flex items-center gap-3">
        <template v-if="!authStore.isAuthenticated">
          <button @click="authStore.login" class="text-sm font-medium text-foreground px-4 py-2 rounded-lg hover:bg-muted transition-colors">
            Sign in
          </button>
          <button class="text-sm font-medium bg-primary text-primary-foreground px-4 py-2 rounded-lg hover:opacity-90 transition-opacity">
            Get started free
          </button>
        </template>
        <template v-else>
          <button class="text-sm font-medium text-foreground px-4 py-2 rounded-lg hover:bg-muted transition-colors">
            Dashboard
          </button>
        </template>
      </div>

      <button
          class="md:hidden p-2 rounded-lg hover:bg-muted transition-colors"
          @click="isOpen = !isOpen"
      >
        <X v-if="isOpen" :size="20" />
        <Menu v-else :size="20" />
      </button>
    </div>

    <div v-if="isOpen" class="md:hidden border-t border-border bg-background px-5 py-4 flex flex-col gap-3">
      <button
          v-for="item in navItems"
          :key="item.name"
          @click="navigateMobile(item.path)"
          class="text-left text-sm font-medium text-muted-foreground py-2 hover:text-foreground transition-colors"
      >
        {{ item.name }}
      </button>

      <div v-if="!authStore.isAuthenticated" class="flex gap-3 pt-2">
        <button @click="authStore.login" class="flex-1 text-sm font-medium border border-border px-4 py-2.5 rounded-lg">
          Sign in
        </button>
        <button class="flex-1 text-sm font-medium bg-primary text-primary-foreground px-4 py-2.5 rounded-lg">
          Get started
        </button>
      </div>
    </div>
  </header>
</template>

