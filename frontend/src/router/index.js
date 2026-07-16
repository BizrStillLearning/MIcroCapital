import { createRouter, createWebHistory } from 'vue-router'
import DashboardLayout from '../components/layouts/DashboardLayout.vue'

import Home from '../views/Home.vue'
import SignIn from '../views/SignIn.vue'
import SignUp from '../views/SignUp.vue'

import Overview from '../views/dashboard/Index.vue'
import Wallet from '../views/dashboard/Wallet.vue'
import Crowdfunding from '../views/dashboard/Crowdfunding.vue'
import Loans from '../views/dashboard/Loans.vue'
import GroupSavings from '../views/dashboard/GroupSavings.vue'
import Settings from '../views/dashboard/Settings.vue'
import Verification from "../views/dashboard/Verification.vue";
import CashManagement from "../views/dashboard/CashManagement.vue";
import AgentOverview from "../views/dashboard/AgentOverview.vue";
import AdminOverview from "../views/dashboard/AdminOverview.vue";
import AdminAgents from "../views/dashboard/AdminAgents.vue";

const isAuthenticated = () => {
    return localStorage.getItem('dummy_token') !== null
}

const routes = [
    { path: '/', component: Home, name: 'Home' },
    { path: '/signin', component: SignIn, name: 'SignIn' },
    { path: '/signup', component: SignUp, name: 'SignUp' },

    {
        path: '/dashboard',
        component: DashboardLayout,
        meta: { requiresAuth: true },
        children: [
            {
                path: '',
                component: Overview,
                name: 'DashboardOverview'
            },
            {
                path: 'agent',
                component: AgentOverview,
                name: 'DashboardAgentOverview'
            },
            {
                path: 'wallet',
                component: Wallet,
                name: 'DashboardWallet'
            },
            {
                path: 'crowdfunding',
                component: Crowdfunding,
                name: 'DashboardCrowdfunding'
            },
            {
                path: 'loans',
                component: Loans,
                name: 'DashboardLoans'
            },
            {
                path: 'community',
                component: GroupSavings,
                name: 'DashboardCommunity'
            },
            {
                path: 'settings',
                component: Settings,
                name: 'DashboardSettings'
            },

            {
                path: 'cash-management',
                component: CashManagement,
                name: 'DashboardCashManagement'
            },
            {
                path: 'verification',
                component: Verification,
                name: 'DashboardVerification'
            },
            {
                path: 'admin',
                component: AdminOverview,
                name: 'DashboardAdminOverview'
            },
            {
                path: 'admin/agents',
                component: AdminAgents,
                name: 'DashboardAdminAgents'
            }
        ]
    }
]

const router = createRouter({
    history: createWebHistory(),
    routes
})

router.beforeEach((to, from, next) => {
    if (to.meta.requiresAuth && !isAuthenticated()) {
        next('/signin')
    } else if ((to.path === '/signin' || to.path === '/signup') && isAuthenticated()) {
        next('/dashboard')
    } else {
        next()
    }
})

export default router


