<script setup>
import { ref, onMounted, onUnmounted } from 'vue'
import { useRouter } from 'vue-router'
import { Menu, X, Shield, Sparkles, Home, Layers } from '@lucide/vue'
import { useAuthStore } from '../stores/auth.js'

const router = useRouter()
const authStore = useAuthStore()
const isOpen = ref(false)
const emit = defineEmits(['nav'])

const activeSection = ref('hero')

const navItems = [
  { name: 'Beranda', path: 'hero', icon: Home },
  { name: 'Cara Kerja', path: 'how-it-works', icon: Layers },
  { name: 'Layanan', path: 'services', icon: Sparkles },
  { name: 'Keamanan', path: 'trust', icon: Shield }
]

const handleScroll = () => {
  const sections = navItems.map(item => item.path)

  for (let i = sections.length - 1; i >= 0; i--) {
    const el = document.getElementById(sections[i])
    if (el) {
      const rect = el.getBoundingClientRect()
      if (rect.top <= 150) {
        activeSection.value = sections[i]
        break
      }
    }
  }
}

onMounted(() => {
  window.addEventListener('scroll', handleScroll)
  handleScroll()
})

onUnmounted(() => {
  window.removeEventListener('scroll', handleScroll)
})

const navigate = (path) => {
  activeSection.value = path
  emit('nav', path)
}

const navigateMobile = (path) => {
  activeSection.value = path
  emit('nav', path)
  isOpen.value = false
}

const navigateToSignIn = () => {
  isOpen.value = false
  router.push('/signin')
}

const navigateToSignUp = () => {
  isOpen.value = false
  router.push('/signup')
}

const navigateToDashboard = () => {
  isOpen.value = false
  router.push('/dashboard')
}
</script>

<template>
  <header class="sticky top-0 z-40 bg-background/90 backdrop-blur-md border-b border-border transition-all">
    <div class="max-w-6xl mx-auto px-5 h-16 flex items-center justify-between">

      <button
          @click="navigate('hero')"
          class="flex items-center gap-2.5 group focus:outline-none"
      >
        <div class="w-8 h-8 rounded-lg bg-primary flex items-center justify-center shadow-sm group-hover:scale-105 transition-transform">
          <span class="text-primary-foreground font-bold text-sm" style="font-family: 'Fraunces', serif;">U</span>
        </div>
        <span class="font-bold text-lg text-foreground tracking-tight" style="font-family: 'Fraunces', serif;">
          Umoja
        </span>
      </button>

      <nav class="hidden md:flex items-center gap-8 text-sm font-medium">
        <button
            v-for="item in navItems"
            :key="item.name"
            @click="navigate(item.path)"
            class="transition-colors relative py-1 after:content-[''] after:absolute after:bottom-0 after:left-0 after:h-0.5 after:bg-primary after:transition-all after:duration-300 hover:text-primary hover:after:w-full"
            :class="activeSection === item.path ? 'text-primary after:w-full' : 'text-muted-foreground after:w-0'"
        >
          {{ item.name }}
        </button>
      </nav>

      <div class="hidden md:flex items-center gap-3">
        <template v-if="!authStore.isAuthenticated">
          <button
              @click="navigateToSignIn"
              class="text-sm font-semibold text-foreground px-4 py-2 rounded-xl hover:bg-muted/80 transition-all"
          >
            Masuk
          </button>
          <button
              @click="navigateToSignUp"
              class="text-sm font-semibold bg-primary text-primary-foreground px-4 py-2 rounded-xl hover:opacity-90 hover:shadow-md hover:shadow-primary/20 transition-all"
          >
            Daftar Gratis
          </button>
        </template>
        <template v-else>
          <button
              @click="navigateToDashboard"
              class="text-sm font-semibold text-foreground px-4 py-2 rounded-xl hover:bg-muted transition-colors border border-border"
          >
            Dasbor
          </button>
        </template>
      </div>

      <button
          class="md:hidden p-2 -mr-2 rounded-xl text-foreground hover:bg-muted/80 active:scale-95 transition-all focus:outline-none"
          @click="isOpen = true"
          aria-label="Buka Menu"
      >
        <Menu :size="24" />
      </button>
    </div>

    <Transition
        enter-active-class="transition-opacity duration-300"
        enter-from-class="opacity-0"
        enter-to-class="opacity-100"
        leave-active-class="transition-opacity duration-300"
        leave-from-class="opacity-100"
        leave-to-class="opacity-0"
    >
      <div
          v-if="isOpen"
          class="fixed inset-0 z-50 bg-black/60 backdrop-blur-sm md:hidden"
          @click="isOpen = false"
      ></div>
    </Transition>

    <Transition
        enter-active-class="transition-transform duration-300 ease-out"
        enter-from-class="translate-x-full"
        enter-to-class="translate-x-0"
        leave-active-class="transition-transform duration-200 ease-in"
        leave-from-class="translate-x-0"
        leave-to-class="translate-x-full"
    >
      <div
          v-if="isOpen"
          class="fixed top-0 right-0 z-[60] w-[85%] max-w-sm h-[100dvh] bg-card border-l border-border flex flex-col md:hidden shadow-2xl overflow-hidden"
      >
        <div class="h-16 px-5 flex items-center justify-between border-b border-border/50 flex-shrink-0">
          <span class="font-bold text-lg text-foreground tracking-tight" style="font-family: 'Fraunces', serif;">Menu</span>
          <button
              @click="isOpen = false"
              class="p-2 -mr-2 rounded-xl bg-muted/50 hover:bg-muted text-foreground transition-colors"
          >
            <X :size="20" />
          </button>
        </div>

        <div class="flex-1 overflow-y-auto px-5 py-6">
          <div class="flex flex-col space-y-2">
            <button
                v-for="item in navItems"
                :key="item.name"
                @click="navigateMobile(item.path)"
                class="flex items-center gap-4 text-left text-base font-medium px-4 py-3.5 rounded-xl hover:bg-muted/60 active:bg-muted transition-all"
                :class="activeSection === item.path ? 'text-primary bg-primary/5' : 'text-foreground'"
            >
              <div
                  v-if="item.icon"
                  class="w-9 h-9 rounded-xl flex items-center justify-center flex-shrink-0 transition-colors"
                  :class="activeSection === item.path ? 'bg-primary text-primary-foreground' : 'bg-primary/10 text-primary'"
              >
                <component :is="item.icon" :size="18" />
              </div>
              <span>{{ item.name }}</span>
            </button>
          </div>
        </div>

        <div class="p-5 border-t border-border/50 bg-muted/10 flex-shrink-0 pb-8">
          <div v-if="!authStore.isAuthenticated" class="flex flex-col gap-3">
            <button
                @click="navigateToSignIn"
                class="w-full text-center text-sm font-semibold border border-border bg-background text-foreground px-4 py-3.5 rounded-xl hover:bg-muted/50 active:scale-[0.99] transition-all shadow-sm"
            >
              Masuk ke Akun
            </button>
            <button
                @click="navigateToSignUp"
                class="w-full text-center text-sm font-semibold bg-primary text-primary-foreground px-4 py-3.5 rounded-xl hover:opacity-90 active:scale-[0.99] transition-all shadow-lg shadow-primary/20"
            >
              Mulai Daftar Gratis
            </button>
          </div>
          <div v-else>
            <button
                @click="navigateToDashboard"
                class="w-full flex items-center justify-center gap-2 text-sm font-semibold bg-primary text-primary-foreground px-4 py-3.5 rounded-xl shadow-md"
            >
              Buka Dasbor
            </button>
          </div>
        </div>

      </div>
    </Transition>
  </header>
</template>

