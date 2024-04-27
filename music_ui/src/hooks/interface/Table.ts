export namespace Table {
    export interface PageAble {
      pageNum: number;
      pageSize: number;
      total: number;
    }
    export interface StateProps {
      tableData: any[];
      pageAble: PageAble;
      searchParam: {
        [key: string]: any;
      };
      searchInitParam: {
        [key: string]: any;
      };
      totalParam: {
        [key: string]: any;
      };
      icon?: {
        [key: string]: any;
      };
    }
  }
  
  export namespace HandleData {
    export type MessageType = "" | "success" | "warning" | "info" | "error";
  }
  
  export namespace Theme {
    export type ThemeType = "light" | "inverted" | "dark";
    export type GreyOrWeakType = "grey" | "weak";
  }
  