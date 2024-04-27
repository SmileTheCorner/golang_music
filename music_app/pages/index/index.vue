<template>
	<view class="content">
		<view class="search">
			<view class="text">
				请输入歌曲名称
			</view>
			<view class="icon">
				<uni-icons type="search" size="30" color="#dcdcdc"></uni-icons>
			</view>
		</view>
		<view class="classify">
			<view class="item">
				推荐
			</view>
			<view class="item">
				歌单
			</view>
			<view class="item">
				流行
			</view>
			<view class="item">
				古风
			</view>
			<view class="item">
				民谣
			</view>
			<view class="item">
				网络
			</view>
			<view class="item">
				古风
			</view>
			<view class="item">
				民谣
			</view>
			<view class="item">
				网络
			</view>
			<view class="item">
				古风
			</view>
			<view class="item">
				民谣
			</view>
		</view>
		<view class="banner">
			<swiper class="swiper" circular :indicator-dots="indicatorDots" :autoplay="autoplay" :interval="interval"
				:duration="duration">
				<swiper-item>
					<image src="https://picsum.photos/400/300"></image>
				</swiper-item>
				<swiper-item>
					<image src="https://picsum.photos/400/300"></image>
				</swiper-item>
				<swiper-item>
					<image src="https://picsum.photos/400/300"></image>
				</swiper-item>
			</swiper>
		</view>
		<view class="ranking">
			<view class="title">排行版</view>
			<view class="scroll">
				<scroll-view class="ranking-scroll" :show-scrollbar="false" scroll-x="true" scroll-left="120">
					<view class="scroll-view-itemA">
						<view class="header">
							<view class="title">
								热歌版
							</view>
							<view class="img">
								<image src="https://picsum.photos/50/50"></image>
							</view>
						</view>
						<view class="content">
							<view class="item">
								<view class="sort">1</view>
								<view class="name">
									<view class="song-name">孤勇者</view>
									<view class="author">陈奕迅</view>
								</view>
							</view>
							<view class="item">
								<view class="sort">2</view>
								<view class="name">
									<view class="song-name">七里香</view>
									<view class="author">周杰伦</view>
								</view>
							</view>
							<view class="item">
								<view class="sort">3</view>
								<view class="name">
									<view class="song-name">三生三辛</view>
									<view class="author">海来啊木</view>
								</view>
							</view>
						</view>
					</view>
					<view class="scroll-view-itemB">
						<view class="header">
							<view class="title">
								热歌版
							</view>
							<view class="img">
								<image src="https://picsum.photos/50/50"></image>
							</view>
						</view>
						<view class="content">
							<view class="item">
								<view class="sort">1</view>
								<view class="name">
									<view class="song-name">孤勇者</view>
									<view class="author">陈奕迅</view>
								</view>
							</view>
							<view class="item">
								<view class="sort">2</view>
								<view class="name">
									<view class="song-name">七里香</view>
									<view class="author">周杰伦</view>
								</view>
							</view>
							<view class="item">
								<view class="sort">3</view>
								<view class="name">
									<view class="song-name">三生三辛</view>
									<view class="author">海来啊木</view>
								</view>
							</view>
						</view>
					</view>
					<view class="scroll-view-itemC">
						<view class="header">
							<view class="title">
								热歌版
							</view>
							<view class="img">
								<image src="https://picsum.photos/50/50"></image>
							</view>
						</view>
						<view class="content">
							<view class="item">
								<view class="sort">1</view>
								<view class="name">
									<view class="song-name">孤勇者</view>
									<view class="author">陈奕迅</view>
								</view>
							</view>
							<view class="item">
								<view class="sort">2</view>
								<view class="name">
									<view class="song-name">七里香</view>
									<view class="author">周杰伦</view>
								</view>
							</view>
							<view class="item">
								<view class="sort">3</view>
								<view class="name">
									<view class="song-name">三生三辛</view>
									<view class="author">海来啊木</view>
								</view>
							</view>
						</view>
					</view>
				</scroll-view>
			</view>
		</view>
		<view class="recommend">
			<view class="title">为你推荐</view>
			<view class="recommend-content">
				<scroll-view class="recommend-scroll" scroll-x="true" scroll-left="120">
					<view class="recommend-item">
						<view class="img">
							<image src="https://picsum.photos/50/50"></image>
							<view class="describe">金典粤语：重温那段流金岁月</view>
						</view>
						<view class="img">
							<image src="https://picsum.photos/50/50"></image>
							<view class="describe">金典粤语：重温那段流金岁月</view>
						</view>
					</view>
					
					<view class="recommend-item">
						<view class="img">
							<image src="https://picsum.photos/50/50"></image>
							<view class="describe">金典粤语：重温那段流金岁月</view>
						</view>
						<view class="img">
							<image src="https://picsum.photos/50/50"></image>
							<view class="describe">金典粤语：重温那段流金岁月</view>
						</view>
					</view>
					
					<view class="recommend-item">
						<view class="img">
							<image src="https://picsum.photos/50/50"></image>
							<view class="describe">金典粤语：重温那段流金岁月</view>
						</view>
						<view class="img">
							<image src="https://picsum.photos/50/50"></image>
							<view class="describe">金典粤语：重温那段流金岁月</view>
						</view>
					</view>
					
					<view class="recommend-item">
						<view class="img">
							<image src="https://picsum.photos/50/50"></image>
							<view class="describe">金典粤语：重温那段流金岁月</view>
						</view>
						<view class="img">
							<image src="https://picsum.photos/50/50"></image>
							<view class="describe">金典粤语：重温那段流金岁月</view>
						</view>
					</view>
					
					<view class="recommend-item">
						<view class="img">
							<image src="https://picsum.photos/50/50"></image>
							<view class="describe">金典粤语：重温那段流金岁月</view>
						</view>
						<view class="img">
							<image src="https://picsum.photos/50/50"></image>
							<view class="describe">金典粤语：重温那段流金岁月</view>
						</view>
					</view>
				</scroll-view>
			</view>
		</view>
	<view class="prefecture">
		<view class="title">
			精彩专区
		</view>
		<view class="prefecture-content">
			<view class="prefecture-item">
				影视专区
			</view>
			<view class="prefecture-item">
				影视专区
			</view>
			<view class="prefecture-item">
				影视专区
			</view>
			<view class="prefecture-item">
				影视专区
			</view>
		</view>
	</view>
	</view>
</template>

<script setup>
	import {
		ref
	} from "vue"
	const indicatorDots = ref(true)
	const autoplay = ref(true)
	const interval = ref(2000)
	const duration = ref(500)
</script>

<style lang="scss" scoped>
	.content {
		display: flex;
		flex-direction: column;
		padding: 25rpx;

		.search {
			box-sizing: border-box;
			width: 100%;
			height: 100rpx;
			border-radius: 50rpx;
			background-color: #ffffff;
			display: flex;

			padding: 0 20rpx;
			flex-direction: row;
			justify-content: space-between;
			align-items: center;

			.text {
				color: #dcdcdc;
			}
		}

		.classify {

			margin: 20rpx 0;
			width: 100%;
			height: 100rpx;
			overflow-x: scroll;
			display: flex;
			flex-direction: row;
			align-items: center;

			.item {
				white-space: nowrap;
				font-size: 40rpx;
				font-weight: bold;
				margin-right: 30rpx;
			}
		}

		.banner {
			.swiper image {
				width: 100%;
				height: 100%;
				object-fit: cover;
			}
		}

		.ranking {
			margin-top: 10rpx;

			.title {
				color: #ffffff;
				font-size: 40rpx;
				font-weight: bold;
				margin-bottom: 10rpx;
			}

			.scroll {
				width: 100%;
				height: 470rpx;
				border-radius: 20rpx;
				overflow: hidden;

				.ranking-scroll {
					box-sizing: border-box;
					width: 100%;
					height: 100%;
					white-space: nowrap;

					.scroll-view-itemA {
						box-sizing: border-box;
						padding: 20rpx;
						height: 100%;
						width: 90%;
						display: inline-block;
						margin-right: 20rpx;
						background: linear-gradient(227.855deg, rgb(255, 180, 172) 20.0094%, rgb(254, 240, 163) 79.1472%);
						border-radius: 20rpx 80.5rpx 20rpx 20rpx;

						.header {
							display: flex;
							flex-direction: row;
							justify-content: space-between;

							.img {
								width: 100rpx;
								height: 100rpx;
								border-radius: 50%;
								overflow: hidden;
							}

							.img image {
								width: 100%;
								height: 100%;
							}
						}

						.content {
							.item {
								display: flex;
								flex-direction: row;

								.sort {
									font-weight: bold;
									font-style: italic;
									font-size: 50rpx;
									margin-right: 20rpx;
								}

								.name {
									.song-name {
										font-weight: bold;
										font-size: 40rpx;
									}

									.author {
										font-size: 30rpx;
										color: rgb(159, 143, 110);
									}
								}
							}
						}

					}

					.scroll-view-itemB {
						box-sizing: border-box;
						padding: 20rpx;
						height: 100%;
						width: 90%;
						display: inline-block;
						margin-right: 20rpx;
						background: linear-gradient(227.855deg, rgb(161, 246, 224) 20.0094%, rgb(239, 254, 210) 79.1472%);
						border-radius: 20rpx 80.5rpx 20rpx 20rpx;

						.header {
							display: flex;
							flex-direction: row;
							justify-content: space-between;

							.img {
								width: 100rpx;
								height: 100rpx;
								border-radius: 50%;
								overflow: hidden;
							}

							.img image {
								width: 100%;
								height: 100%;
							}
						}

						.content {
							.item {
								display: flex;
								flex-direction: row;

								.sort {
									font-weight: bold;
									font-style: italic;
									font-size: 50rpx;
									margin-right: 20rpx;
								}

								.name {
									.song-name {
										font-weight: bold;
										font-size: 40rpx;
									}

									.author {
										font-size: 30rpx;
										color: rgb(159, 143, 110);
									}
								}
							}
						}

					}

					.scroll-view-itemC {
						box-sizing: border-box;
						padding: 20rpx;
						height: 100%;
						width: 90%;
						display: inline-block;
						background: linear-gradient(227.855deg, rgb(255, 180, 172) 20.0094%, rgb(254, 240, 163) 79.1472%);
						border-radius: 20rpx 80.5rpx 20rpx 20rpx;

						.header {
							display: flex;
							flex-direction: row;
							justify-content: space-between;

							.img {
								width: 100rpx;
								height: 100rpx;
								border-radius: 50%;
								overflow: hidden;
							}

							.img image {
								width: 100%;
								height: 100%;
							}
						}

						.content {
							.item {
								display: flex;
								flex-direction: row;

								.sort {
									font-weight: bold;
									font-style: italic;
									font-size: 50rpx;
									margin-right: 20rpx;
								}

								.name {
									.song-name {
										font-weight: bold;
										font-size: 40rpx;
									}

									.author {
										font-size: 30rpx;
										color: rgb(159, 143, 110);
									}
								}
							}
						}
					}
				}
			}
		}

		.recommend {
			.title {
				color: #ffffff;
				font-size: 40rpx;
				font-weight: bold;
				margin: 10rpx 0;
			}

			.recommend-content {
				box-sizing: border-box;
				width: 100%;
				.recommend-scroll {
					width: 100%;
					height: 100%;
					box-sizing: border-box;
					white-space: nowrap;
					.recommend-item {
						display: inline-block;
						margin: 0 20rpx 0 0;
						.img {
							.describe {
								white-space: pre-wrap;
								width: 200rpx;
								font-size: 30rpx;
								overflow: hidden;
								display: -webkit-box;
								-webkit-box-orient: vertical;
								-webkit-line-clamp: 2;
							}
						}
						.img image {
							width: 200rpx;
							height: 200rpx;
							border-radius: 10rpx;
						}

					}
					.recommend-item:last-child{
						margin-right: 0;
					}
				}
			}

		}
	.prefecture{
		.title{
			color: #ffffff;
			font-size: 40rpx;
			font-weight: bold;
		}
		.prefecture-content{
			width: 100%;
			display: grid;
			grid-template-columns: 1fr 1fr;
			gap: 20rpx;
			overflow-x: scroll;
			.prefecture-item{
				border-radius: 10rpx;
				padding: 20rpx;
				color: #ffffff;
				font-size: 30rpx;
				font-weight: bold;
				width: 400rpx;
				height: 200rpx;
				background-color: red;
				background-image: url('https://picsum.photos/200/100');
			}
		}
	}
	}
</style>