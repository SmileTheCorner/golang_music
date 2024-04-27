import { defineComponent, ref } from "vue";

export default defineComponent({
  name: "useDialog",
  props: {
    dialogConfig: {
      type: Object as () => Record<string, any>,
      default: () => ({}),
    },
  },
  setup(props, ctx) {
    const { slots } = ctx;
    console.log("dialogConfig", props.dialogConfig);
    return () => (
      <>
        <div class="dialogContainer">
          <el-dialog
            v-model={props.dialogConfig.visible}
            {...props.dialogConfig}
            v-slots={{
              footer: () => <>{slots.footer?.()}</>,
            }}
          >
            <div class="dialogContent">{slots.default?.()}</div>
          </el-dialog>
        </div>
      </>
    );
  },
});
