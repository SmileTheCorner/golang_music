import { defineComponent, PropType } from "vue";
import { FormOption, FormItem } from "./interface/index";
import "./index.scss";

export default defineComponent({
  //组件名称
  name: "ProForm",
  //接收父组件传递过来的值
  props: {
    formOption: {
      type: Object as PropType<FormOption>,
      required: true,
      default: () => {},
    },
  },
  setup(props, context) {
    const { slots } = context;
    return () => (
      <>
        <div class="elFormContainer">
          <el-form
            disabled={props.formOption.disabled}
            inline={props.formOption.inline}
            label-position={props.formOption.labelPosition}
            label-width={props.formOption.labelWidth}
          >
            <div class="formContainer">
              {props.formOption.formItems.map(
                (item: FormItem, index: number) => {
                  {
                    /* 普通输入框 */
                  }
                  if (item.type === "input") {
                    return (
                      <el-form-item
                        class="formItem"
                        label={item.label}
                        prop={item.prop}
                        key={index}
                      >
                        <el-input
                          v-model={item.value}
                          placeholder={item.placeholder}
                        ></el-input>
                      </el-form-item>
                    );
                  }

                  {
                    /* 下拉选择框 */
                  }
                  if (item.type === "select") {
                    return (
                      <el-form-item
                        class="formItem"
                        label={item.label}
                        prop={item.prop}
                        key={index}
                      >
                        <el-select
                          v-model={item.value}
                          placeholder={item.placeholder}
                        >
                          {item.options?.map((option: any, index: number) => {
                            return (
                              <el-option
                                label={option.label}
                                value={option.value}
                                key={index}
                              ></el-option>
                            );
                          })}
                        </el-select>
                      </el-form-item>
                    );
                  }

                  {
                    /* {单选框} */
                  }
                  if (item.type === "radio") {
                    return (
                      <el-form-item
                        class="formItem"
                        label={item.label}
                        prop={item.prop}
                        key={index}
                      >
                        <el-radio-group v-model={item.value} size="small">
                          {item.options?.map((option: any, index: number) => {
                            return (
                              <el-radio
                                label={option.label}
                                value={option.value}
                                key={index}
                              ></el-radio>
                            );
                          })}
                        </el-radio-group>
                      </el-form-item>
                    );
                  }

                  {
                    /* {多选框} */
                  }
                  if (item.type === "checkbox") {
                    return (
                      <el-form-item
                        class="formItem"
                        label={item.label}
                        prop={item.prop}
                        key={index}
                      >
                        <el-checkbox-group v-model={item.value} size="small">
                          {item.options?.map((option: any, index: number) => {
                            return (
                              <el-checkbox
                                label={option.label}
                                value={option.value}
                                key={index}
                              ></el-checkbox>
                            );
                          })}
                        </el-checkbox-group>
                      </el-form-item>
                    );
                  }

                  {
                    /* {switch 开关选择器} */
                  }
                  if (item.type === "switch") {
                    return (
                      <el-form-item
                        class="formItem"
                        label={item.label}
                        prop={item.prop}
                        key={index}
                      >
                        <el-switch v-model={item.value}></el-switch>
                      </el-form-item>
                    );
                  }

                  {
                    /* {el-time-picker 时间选择器} */
                  }
                  if (item.type === "timepicker") {
                    return (
                      <el-form-item
                        class="formItem"
                        type={item.dateTimeType}
                        label={item.label}
                        prop={item.prop}
                        key={index}
                      >
                        <el-time-picker
                          v-model={item.value}
                          placeholder={item.placeholder}
                          is-range={item.isRange}
                          range-separator="至"
                          start-placeholder="开始时间"
                          end-placeholder="结束时间"
                        ></el-time-picker>
                      </el-form-item>
                    );
                  }

                  {
                    /* {el-date-picker 日期选择器} */
                  }
                  if (item.type === "datepicker") {
                    return (
                      <el-form-item
                        class="formItem"
                        label={item.label}
                        prop={item.prop}
                        key={index}
                      >
                        <el-date-picker
                          type={item.dateTimeType}
                          range-separator="至"
                          start-placeholder="开始时间"
                          end-placeholder="结束时间"
                          v-model={item.value}
                          placeholder={item.placeholder}
                        ></el-date-picker>
                      </el-form-item>
                    );
                  }

                  {
                    /* {textarea 文本框输入} */
                  }
                  if (item.type === "textarea") {
                    return (
                      <el-form-item
                        class="formItem"
                        label={item.label}
                        prop={item.prop}
                        key={index}
                      >
                        <el-input
                          type="textarea"
                          v-model={item.value}
                          placeholder={item.placeholder}
                        ></el-input>
                      </el-form-item>
                    );
                  }
                }
              )}
              {slots.default ? (
                <el-form-item class="formItem operate">
                  {slots.default?.()}
                </el-form-item>
              ) : null}
            </div>
          </el-form>
        </div>
      </>
    );
  },
});
