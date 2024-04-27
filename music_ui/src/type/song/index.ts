export interface Song {
  createBy: string;
  createdAt: string;
  deletedAt: string;
  id: string;
  isNew: boolean;
  lyricsId: string;
  playCount: number;
  rankingListId: string;
  recommend: boolean;
  songCover: string;
  songDescription: string;
  songName: string;
  songSingerId: string;
  songType: string;
  songUrl: string;
  updateBy: string;
  updatedAt: string;
}

export interface SingerList {
	id: number,
	name: string,
	classify?:number,
	children:[]
}
