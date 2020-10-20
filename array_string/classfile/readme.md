# ClassReader

> 只是一个方便读取数据的工具，仅此而已

# ClassFile

> 解析类数据的入口
>
> 依次解析：
>
> - **magic**
> - **minor_version**
> - **major_version**
> - resolveConstantPool
>   - **constant_pool_count**
>   - **constant_pool[constant_pool_count - 1]**
>     - 循环创建以下信息
>     - `ConstantIntegerInfo`
>     - `ConstantFloatInfo`
>     - `ConstantLongInfo`
>     - `ConstantDoubleInfo`
>     - `ConstantUtf8Info`
>     - `ConstantStringInfo`
>     - `ConstantClassInfo`
>     - `ConstantFieldrefInfo`
>     - `ConstantMethodrefInfo`
>     - `ConstantInterfaceMethodrefInfo`
>     - `ConstantNameAndTypeInfo`
>     - `ConstantMethodTypeInfo`
>     - `ConstantMethodHandleInfo`
>     - `ConstantInvokeDynamicInfo`
> - **access_flags**
> - **this_class**
> - **super_class**
> - readUint16s
>   - **interfaces_count**
>   - **interfaces[interfaces_count]**
> - resolveMembers
>   - **fileds_count**
>   - **fileds[fields_count]**
> - resolveMembers
>   - **methods_count**
>   - **methods[methods_count]**
> - resolveAttributes
>   - **attributes_count**
>   - **attributes[attributes_count]**
>     - 循环创建以下信息
>     - `BootstrapMethodsAttribute`
>     - `CodeAttribute`
>     - `ConstantValueAttribute`
>     - `EnclosingMethodAttribute`
>     - `ExceptionsAttribute`
>     - `InnerClassesAttribute`
>     - `LineNumberTableAttribute`
>     - `LocalVariableTableAttribute`
>     - `LocalVariableTypeTableAttribute`
>     - `SignatureAttribute`
>     - `SourceFileAttribute`
>     - `UnparsedAttribute`

