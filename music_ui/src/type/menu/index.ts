declare namespace Menu {
    interface MenuOptions {
      path?: string,
      name?: string,
      component?: string | (() => Promise<unknown>),
      redirect?: string,
      meta?: MetaProps,
      children?: MenuOptions[],
    }
    interface MetaProps {
      id?:string,
      parentId?:string,
      icon?: string,
      title?: string,
      menuType?: string,
      isFrame?: number,
      isCache?: number,	
      visible?: number,
      status?: number,
      perms?: string,
      orderNum?:number,
      query?: string;
      isFull?: number,
      remark?: string,
      isLink?: string, 
    }
  }