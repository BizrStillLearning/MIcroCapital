import {
    Home,
    CircleDollarSign,
    Users,
    Settings,
    Landmark,
    ShieldCheck,
    BarChart3,
    UsersRound
} from '@lucide/vue'

export const getAsideMenu = (role) => {
    const dashboardHome = { name: 'Ringkasan', path: '/dashboard', icon: Home }
    const dashboardSettings = { name: 'Pengaturan', path: '/dashboard/settings', icon: Settings }

    if (role === 'member') {
        return [
            dashboardHome,
            { name: 'Dompet Kas', path: '/dashboard/wallet', icon: CircleDollarSign },
            { name: 'Urun Dana', path: '/dashboard/crowdfunding', icon: Users },
            { name: 'Pinjaman', path: '/dashboard/loans', icon: CircleDollarSign },
            dashboardSettings
        ]
    }

    if (role === 'agent') {
        return [
            dashboardHome,
            { name: 'Manajemen Tunai', path: '/dashboard/cash-management', icon: Landmark },
            { name: 'Verifikasi Warga', path: '/dashboard/verification', icon: ShieldCheck },
            dashboardSettings
        ]
    }

    if (role === 'admin') {
        return [
            { name: 'Analitik Global', path: '/dashboard/admin', icon: BarChart3 },
            { name: 'Manajemen Agen', path: '/dashboard/admin/agents', icon: UsersRound },
            dashboardSettings
        ]
    }

    return [dashboardHome, dashboardSettings]
}

