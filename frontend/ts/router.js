import {createRouter,createWebHistory} from "vue-router";
import AdminPage from "../admin/page.vue";
import LoginPage from "../login/main.vue";
import ResHelper from "../admin/res.vue";
import UserHelper from "../admin/user.vue";
import MainHelper from "../admin/main.vue";
import MaintainHelper from "../admin/maintain.vue";
import ReportHelper from "../admin/report.vue";

const routers = createRouter({
    history: createWebHistory(),
    routes: [
        {
            path: '/',
            name:"index",
            redirect: '/login'
        },
        {
          path:'/login',
            name:"login",
            component:LoginPage
        },
        {
            path: "/admin",
            name: "admin",
            component: AdminPage,
            children:[
                {path:"main", component:MainHelper},
                {path: "resadmin", component: ResHelper,},
                {path:"useradmin",component:UserHelper},
                {path:"maintainadmin",component:MaintainHelper},
                {path:"reportadmin",component:ReportHelper}
            ]
        },
    ]
    ,});


export default routers;
