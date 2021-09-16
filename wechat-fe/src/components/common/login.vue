<template>
    <div>
        <el-input class="loginInput" style="width: 55%" v-model="input" placeholder="请输入昵称"></el-input>
		<el-button style="margin-left: 37%;" size="small" type="success" @click="login">点击登陆</el-button>
	</div>
</template>
<script>
import http from '../../assets/http.js'
export default {
    data() {
            return {
                input: "",
            }
        },
	methods: {
		ready() {
			http.get("/wechat/ready").then((response) => {
				console.log(response);
			})
		},
		login() {
			let params = {
				nickName: 	this.input,
				loginTime: 	new Date().getTime().toString()
			}
			console.log(params)
			http.post("/wechat/login", params).then((res) => {
				console.log(res)
			})
		}
	},
}
</script>

<style>
.loginInput {
	position: relative;
	margin-left: 20%;
	margin-top: 50%;
	margin-bottom: 40px;
}
</style>