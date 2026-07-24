<script setup>
import { reactive, ref } from 'vue'
import { useRouter } from 'vue-router'
import { ArrowLeft, ArrowRight, Smartphone, Lock } from '@lucide/vue'
import { useAuthStore } from '../stores/authStore.js'

const router = useRouter()
const authStore = useAuthStore()

const form = reactive({
  phone: '',
  pin: ''
})

const errorMessage = ref('')

const handleSignIn = async () => {
  errorMessage.value = ''

  if (!form.phone || !form.pin) {
    errorMessage.value = 'Nomor telepon dan PIN wajib diisi'
    return
  }

  const success = await authStore.login(form.phone, form.pin)

  if (success) {
    const role = authStore.userRole
    if (role === 'admin') {
      router.push('/dashboard/admin')
    } else if (role === 'agent') {
      router.push('/dashboard/agent')
    } else {
      router.push('/dashboard')
    }
  } else {
    errorMessage.value = authStore.error
  }
}
</script>

<template>
  <div class="min-h-screen bg-background flex" style="font-family: 'DM Sans', sans-serif;">

    <div class="hidden md:flex md:w-1/2 relative bg-primary items-center justify-center overflow-hidden">
      <div class="absolute inset-0">
        <img src="https://images.unsplash.com/photo-1542884748-2b87b36f6b90?w=1200&h=1080&fit=crop&auto=format" alt="Community" class="w-full h-full object-cover opacity-30 mix-blend-multiply" />
        <div class="absolute inset-0 bg-gradient-to-t from-primary via-primary/80 to-transparent"></div>
      </div>
      <div class="relative z-10 p-12 max-w-lg text-primary-foreground">
        <div class="flex items-center gap-2.5 mb-8">
          <div class="w-10 h-10 rounded-lg bg-white flex items-center justify-center">
            <span class="text-primary font-bold text-xl" style="font-family: 'Fraunces', serif;">U</span>
          </div>
          <span class="font-bold text-2xl text-white" style="font-family: 'Fraunces', serif;">Umoja</span>
        </div>
        <h1 class="text-4xl lg:text-5xl font-bold mb-6 leading-tight" style="font-family: 'Fraunces', serif;">Selamat Datang Kembali.</h1>
        <p class="text-lg opacity-80 leading-relaxed">Masuk untuk mengelola dana Anda dan terhubung kembali dengan komunitas.</p>
      </div>
    </div>

    <div class="w-full md:w-1/2 flex items-center justify-center p-5 sm:p-12 relative">
      <div class="w-full max-w-md">
        <button @click="router.push('/')" class="absolute top-8 left-5 sm:left-12 text-sm text-muted-foreground hover:text-foreground inline-flex items-center gap-2 transition-colors">
          <ArrowLeft :size="16" /> Kembali
        </button>

        <div class="mb-10 mt-8 md:mt-0">
          <h2 class="text-3xl font-bold text-foreground mb-2" style="font-family: 'Fraunces', serif;">Masuk ke Akun</h2>
          <p class="text-muted-foreground text-sm">Masukkan nomor telepon dan PIN Anda.</p>
        </div>

        <div v-if="errorMessage" class="bg-red-50 text-red-600 p-4 rounded-xl mb-6 text-sm font-bold border border-red-200">
          {{ errorMessage }}
        </div>

        <form @submit.prevent="handleSignIn" class="space-y-5">
          <div class="space-y-1.5">
            <label class="text-sm font-semibold text-foreground" style="font-family: 'DM Mono', monospace;">Nomor Telepon</label>
            <div class="relative">
              <Smartphone class="absolute left-3.5 top-1/2 -translate-y-1/2 text-muted-foreground" :size="18" />
              <input v-model="form.phone" type="tel" placeholder="081234567890" class="w-full bg-card border border-border rounded-xl py-3 pl-10 pr-4 text-sm focus:outline-none focus:border-primary focus:ring-1 focus:ring-primary transition-all" required />
            </div>
          </div>

          <div class="space-y-1.5">
            <div class="flex items-center justify-between">
              <label class="text-sm font-semibold text-foreground" style="font-family: 'DM Mono', monospace;">PIN 4-Digit</label>
              <button type="button" class="text-xs text-primary font-semibold hover:underline">Lupa PIN?</button>
            </div>
            <div class="relative">
              <Lock class="absolute left-3.5 top-1/2 -translate-y-1/2 text-muted-foreground" :size="18" />
              <input v-model="form.pin" type="password" inputmode="numeric" maxlength="4" placeholder="••••" class="w-full bg-card border border-border rounded-xl py-3 pl-10 pr-4 text-sm tracking-widest focus:outline-none focus:border-primary focus:ring-1 focus:ring-primary transition-all" required />
            </div>
          </div>

          <button type="submit" :disabled="authStore.isLoading" class="w-full bg-primary text-primary-foreground font-semibold py-3.5 rounded-xl hover:opacity-90 hover:shadow-lg transition-all flex items-center justify-center gap-2 mt-2 disabled:opacity-50 disabled:cursor-not-allowed">
            {{ authStore.isLoading ? 'Memproses...' : 'Masuk' }} <ArrowRight v-if="!authStore.isLoading" :size="18" />
          </button>
        </form>

        <div class="mt-8 text-center text-sm text-muted-foreground">
          Belum punya akun?
          <button @click="router.push('/signup')" class="text-primary font-bold hover:underline ml-1" style="font-family: 'DM Mono', monospace;">Daftar gratis</button>
        </div>
      </div>
    </div>
  </div>
</template>



