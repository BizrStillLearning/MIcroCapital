<script setup>
import { Menu, PanelLeftClose, PanelLeftOpen, Bell } from '@lucide/vue'
import { computed } from 'vue'
import { useAuthStore } from "../../stores/authStore.js"

const authStore = useAuthStore()
const user = computed(() => authStore.user || {})

const props = defineProps({
  isCollapsed: Boolean
})

const emit = defineEmits(['toggleCollapse', 'toggleMobile'])

const getInitials = (name) => {
  if (!name) return 'U'
  const names = name.split(' ')
  if (names.length >= 2) {
    return (names[0][0] + names[1][0]).toUpperCase()
  }
  return name.substring(0, 2).toUpperCase()
}
</script>

<template>
  <header class="h-16 bg-card/80 backdrop-blur-md border-b border-border flex items-center justify-between px-4 sm:px-6 sticky top-0 z-10">

    <div class="flex items-center gap-3">
      <button
          @click="emit('toggleMobile')"
          class="md:hidden p-2 -ml-2 text-muted-foreground hover:text-foreground hover:bg-muted rounded-xl transition-colors"
      >
        <Menu :size="22" />
      </button>

      <button
          @click="emit('toggleCollapse')"
          class="hidden md:flex p-2 -ml-2 text-muted-foreground hover:text-foreground hover:bg-muted rounded-xl transition-colors"
      >
        <PanelLeftOpen v-if="isCollapsed" :size="20" />
        <PanelLeftClose v-else :size="20" />
      </button>

    </div>

    <div class="flex items-center gap-4">
      <button class="relative p-2 text-muted-foreground hover:text-foreground rounded-full hover:bg-muted transition-colors">
        <Bell :size="18" />
        <span class="absolute top-1.5 right-1.5 w-2 h-2 bg-accent rounded-full border border-card"></span>
      </button>

      <div class="flex items-center gap-3 cursor-pointer p-1 pr-2 rounded-full hover:bg-muted transition-colors border border-transparent hover:border-border">
        <div class="w-10 h-10 bg-primary/10 text-primary font-bold rounded-full flex items-center justify-center">
          {{ getInitials(user.name) }}
        </div>
        <div class="hidden sm:block text-left">
          <p class="text-sm font-bold text-foreground">{{ user.name || 'Pengguna' }}</p>
          <p class="text-xs text-muted-foreground capitalize">{{ user.role || 'Member' }}</p>
        </div>
      </div>
    </div>
  </header>
</template>

