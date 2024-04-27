<template>
	<div class="login-container">
		<div class="login-box">
			<div class="login-left">
				<img src="@/assets/images/login_left.gif" alt="login" />
			</div>
			<div class="login-form">
				<div class="login-logo">
					<img class="login-icon" src="@/assets/images/music.png" alt="" />
					<h2 class="logo-text">月报系统</h2>
				</div>
				<LoginForm />
			</div>
		</div>
	</div>
</template>

<script setup lang="ts" name="login">
import LoginForm from "./components/LoginForm.vue";
</script>

<style scoped lang="scss">
@import "./index.scss"
</style>
