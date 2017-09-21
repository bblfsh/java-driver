| Path | Action |
|------|--------|
| /self::\*\[@InternalType='CompilationUnit'\] | File |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='QualifiedName'\] | Expression, Identifier, Qualified |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='SimpleName'\] | Expression, Identifier |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[\(@InternalType='MethodDeclaration'\) or \(@InternalType='TypeDeclaration'\)\]/self::\*\[child::\(@InternalType='Modifier'\) and \(@Token='public'\)\] | Visibility, World |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[\(@InternalType='MethodDeclaration'\) or \(@InternalType='TypeDeclaration'\)\]/self::\*\[child::\(@InternalType='Modifier'\) and \(@Token='private'\)\] | Visibility, Type |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[\(@InternalType='MethodDeclaration'\) or \(@InternalType='TypeDeclaration'\)\]/self::\*\[child::\(@InternalType='Modifier'\) and \(@Token='protected'\)\] | Visibility, Subtype |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[\(@InternalType='MethodDeclaration'\) or \(@InternalType='TypeDeclaration'\)\]/self::\*\[not\(child::\(@InternalType='Modifier'\) and \(\(@Token='public'\) or \(@Token='private'\) or \(@Token='protected'\)\)\)\] | Visibility, Package |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='PackageDeclaration'\] | Declaration, Package |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='ImportDeclaration'\] | Declaration, Import |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='ImportDeclaration'\]/\*\[@InternalType='QualifiedName'\] | Pathname, Import |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='AnonymousClassDeclaration'\] | Expression, Declaration, Type, Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='AnonymousClassDeclaration'\]/\*\[@internalRole\]\[@internalRole='bodyDeclarations'\] | Body |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='AnnotationTypeDeclaration'\] | Declaration, Type, Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='AnnotationTypeDeclaration'\]/\*\[@internalRole\]\[@internalRole='bodyDeclarations'\] | Body |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='EnumDeclaration'\] | Declaration, Type, Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='TypeDeclaration'\] | Declaration, Type |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='TypeDeclarationStatement'\] | Statement, Declaration, Type, Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='MethodDeclaration'\] | Declaration, Function |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='MethodDeclaration'\]/\*\[@internalRole\]\[@internalRole='name'\] | Function, Name |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='MethodDeclaration'\]/\*\[@internalRole\]\[@internalRole='body'\] | Function, Body |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='MethodDeclaration'\]/\*\[@internalRole\]\[@internalRole='parameters'\] | Function, Argument |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='MethodDeclaration'\]/\*\[@internalRole\]\[@internalRole='parameters'\]/self::\*\[@varargs\]\[@varargs='true'\] | Function, ArgsList |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='MethodDeclaration'\]/\*\[@internalRole\]\[@internalRole='parameters'\]/\*\[@internalRole\]\[@internalRole='name'\] | Function, Name |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='LambdaExpression'\] | Declaration, Function, Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='LambdaExpression'\]/\*\[@internalRole\]\[@internalRole='body'\] | Function, Body |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='LambdaExpression'\]/\*\[@internalRole\]\[@internalRole='parameters'\] | Function, Argument |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='LambdaExpression'\]/\*\[@internalRole\]\[@internalRole='parameters'\]/self::\*\[@varargs\]\[@varargs='true'\] | Function, ArgsList |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='LambdaExpression'\]/\*\[@internalRole\]\[@internalRole='parameters'\]/\*\[@internalRole\]\[@internalRole='name'\] | Function, Name |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='AnnotationTypeMemberDeclaration'\] | Declaration, Type, Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='EnumConstantDeclaration'\] | Declaration, Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='FieldDeclaration'\] | Declaration, Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='Initializer'\] | Initialization, Block, Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='SingleVariableDeclaration'\] | Declaration, Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='VariableDeclarationExpression'\] | Expression, Declaration, Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='VariableDeclarationFragment'\] | Declaration, Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='VariableDeclarationStatement'\] | Statement, Declaration, Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='BooleanLiteral'\] | Expression, Literal, Boolean |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='CharacterLiteral'\] | Expression, Literal, Character |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='NullLiteral'\] | Expression, Literal, Null |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='NumberLiteral'\] | Expression, Literal, Number |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='StringLiteral'\] | Expression, Literal, String |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='TypeLiteral'\] | Expression, Literal, Type |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='ClassInstanceCreation'\] | Expression, Call, Instance |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='ClassInstanceCreation'\]/\*\[@internalRole\]\[@internalRole='type'\] | Call, Callee |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='ClassInstanceCreation'\]/\*\[@internalRole\]\[@internalRole='arguments'\] | Call, Argument, Positional |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='ConstructorInvocation'\] | Statement, Call, Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='ConstructorInvocation'\]/\*\[@internalRole\]\[@internalRole='type'\] | Call, Callee |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='ConstructorInvocation'\]/\*\[@internalRole\]\[@internalRole='arguments'\] | Call, Argument, Positional |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='MethodInvocation'\] | Expression, Call |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='MethodInvocation'\]/\*\[@internalRole\]\[@internalRole='expression'\] | Call, Receiver |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='MethodInvocation'\]/\*\[@internalRole\]\[@internalRole='name'\] | Call, Callee |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='MethodInvocation'\]/\*\[@internalRole\]\[@internalRole='arguments'\] | Call, Argument, Positional |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='SuperConstructorInvocation'\] | Statement, Call, Base, Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='SuperConstructorInvocation'\]/\*\[@internalRole\]\[@internalRole='expression'\] | Call, Receiver |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='SuperConstructorInvocation'\]/\*\[@internalRole\]\[@internalRole='arguments'\] | Call, Argument, Positional |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='SuperMethodInvocation'\] | Expression, Call, Base |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='SuperMethodInvocation'\]/\*\[@internalRole\]\[@internalRole='qualifier'\] | Call, Callee |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='SuperMethodInvocation'\]/\*\[@internalRole\]\[@internalRole='name'\] | Call, Callee |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='SuperMethodInvocation'\]/\*\[@internalRole\]\[@internalRole='arguments'\] | Call, Argument, Positional |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='IfStatement'\] | Statement, If |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='IfStatement'\]/\*\[@internalRole\]\[@internalRole='expression'\] | If, Condition |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='IfStatement'\]/\*\[@internalRole\]\[@internalRole='thenStatement'\] | If, Then, Body |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='IfStatement'\]/\*\[@internalRole\]\[@internalRole='elseStatement'\] | If, Else, Body |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='SwitchStatement'\] | Statement, Switch |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='SwitchStatement'\]/\*\[@internalRole\]\[@internalRole='expression'\] | Expression, Switch |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='SwitchStatement'\]/\*\[@InternalType='SwitchCase'\] | Statement, Switch |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='SwitchStatement'\]/\*\[@InternalType='SwitchCase'\]/self::\*\[child::\*\] | Case |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='SwitchStatement'\]/\*\[@InternalType='SwitchCase'\]/self::\*\[child::\*\]/\*\[@internalRole\]\[@internalRole='expression'\] | Expression, Switch, Case, Condition |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='SwitchStatement'\]/\*\[@InternalType='SwitchCase'\]/self::\*\[not\(child::\*\)\] | Default |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='SwitchStatement'\]/\*\[@InternalType='ExpressionStatement'\] | Switch, Case, Body |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='EnhancedForStatement'\] | Statement, For, Iterator |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='EnhancedForStatement'\]/\*\[@internalRole\]\[@internalRole='parameter'\] | For, Iterator |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='EnhancedForStatement'\]/\*\[@internalRole\]\[@internalRole='expression'\] | Expression, For |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='EnhancedForStatement'\]/\*\[@internalRole\]\[@internalRole='body'\] | For, Body |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='ForStatement'\] | Statement, For |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='ForStatement'\]/\*\[@internalRole\]\[@internalRole='initializers'\] | For, Initialization |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='ForStatement'\]/\*\[@internalRole\]\[@internalRole='expression'\] | Expression, For, Condition |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='ForStatement'\]/\*\[@internalRole\]\[@internalRole='updaters'\] | For, Update |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='ForStatement'\]/\*\[@internalRole\]\[@internalRole='body'\] | For, Body |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='WhileStatement'\] | Statement, While |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='WhileStatement'\]/\*\[@internalRole\]\[@internalRole='expression'\] | Expression, While, Condition |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='WhileStatement'\]/\*\[@internalRole\]\[@internalRole='body'\] | While, Body |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='DoStatement'\] | Statement, DoWhile |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='DoStatement'\]/\*\[@internalRole\]\[@internalRole='expression'\] | DoWhile, Condition |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='DoStatement'\]/\*\[@internalRole\]\[@internalRole='body'\] | DoWhile, Body |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='InfixExpression'\] | Expression, Binary, Operator |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='InfixExpression'\]/self::\*\[@operator\]\[@operator='\+'\] | Add |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='InfixExpression'\]/self::\*\[@operator\]\[@operator='\-'\] | Substract |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='InfixExpression'\]/self::\*\[@operator\]\[@operator='\*'\] | Multiply |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='InfixExpression'\]/self::\*\[@operator\]\[@operator='/'\] | Divide |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='InfixExpression'\]/self::\*\[@operator\]\[@operator='%'\] | Modulo |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='InfixExpression'\]/self::\*\[@operator\]\[@operator='<<'\] | Bitwise, LeftShift |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='InfixExpression'\]/self::\*\[@operator\]\[@operator='>>'\] | Bitwise, RightShift |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='InfixExpression'\]/self::\*\[@operator\]\[@operator='>>>'\] | Bitwise, RightShift, Unsigned |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='InfixExpression'\]/self::\*\[@operator\]\[@operator='&'\] | Bitwise, And |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='InfixExpression'\]/self::\*\[@operator\]\[@operator='\|'\] | Bitwise, Or |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='InfixExpression'\]/self::\*\[@operator\]\[@operator='&&'\] | Boolean, And |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='InfixExpression'\]/self::\*\[@operator\]\[@operator='\|\|'\] | Boolean, Or |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='InfixExpression'\]/self::\*\[@operator\]\[@operator='^'\] | Boolean, Xor |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='InfixExpression'\]/\*\[@internalRole\]\[@internalRole='leftOperand'\] | Expression, Binary, Left |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='InfixExpression'\]/\*\[@internalRole\]\[@internalRole='rightOperand'\] | Expression, Binary, Right |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='PostfixExpression'\] | Expression, Operator, Unary, Postfix |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='PostfixExpression'\]/self::\*\[@operator\]\[@operator='\+\+'\] | Increment |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='PostfixExpression'\]/self::\*\[@operator\]\[@operator='\-\-'\] | Increment |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='PrefixExpression'\] | Expression, Operator, Unary |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='PrefixExpression'\]/self::\*\[@operator\]\[@operator='\+\+'\] | Increment |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='PrefixExpression'\]/self::\*\[@operator\]\[@operator='\-\-'\] | Decrement |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='PrefixExpression'\]/self::\*\[@operator\]\[@operator='\+'\] | Positive |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='PrefixExpression'\]/self::\*\[@operator\]\[@operator='\-'\] | Negative |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='PrefixExpression'\]/self::\*\[@operator\]\[@operator='~'\] | Bitwise, Not |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='PrefixExpression'\]/self::\*\[@operator\]\[@operator='\!'\] | Boolean, Not |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='Assignment'\] | Expression, Assignment, Operator, Binary |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='Assignment'\]/\*\[@internalRole\]\[@internalRole='leftHandSide'\] | Assignment, Binary, Left |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='Assignment'\]/\*\[@internalRole\]\[@internalRole='rightHandSide'\] | Assignment, Binary, Right |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='Assignment'\]/self::\*\[@operator\]\[@operator='\+='\] | Add |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='Assignment'\]/self::\*\[@operator\]\[@operator='\-='\] | Substract |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='Assignment'\]/self::\*\[@operator\]\[@operator='\*='\] | Multiply |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='Assignment'\]/self::\*\[@operator\]\[@operator='/='\] | Divide |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='Assignment'\]/self::\*\[@operator\]\[@operator='%='\] | Modulo |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='Assignment'\]/self::\*\[@operator\]\[@operator='&='\] | Bitwise, And |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='Assignment'\]/self::\*\[@operator\]\[@operator='\|='\] | Bitwise, Or |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='Assignment'\]/self::\*\[@operator\]\[@operator='^='\] | Bitwise, Xor |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='Assignment'\]/self::\*\[@operator\]\[@operator='<<='\] | Bitwise, LeftShift |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='Assignment'\]/self::\*\[@operator\]\[@operator='>>='\] | Bitwise, RightShift |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='Assignment'\]/self::\*\[@operator\]\[@operator='>>>='\] | Bitwise, RightShift, Unsigned |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='ArrayType'\] | Type, Primitive, List |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='IntersectionType'\] | Type, And |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='NameQualifiedType'\] | Type, Name, Qualified |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='ParameterizedType'\] | Type, Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='PrimitiveType'\] | Type, Primitive |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='QualifiedType'\] | Type, Qualified |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='SimpleType'\] | Type |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='UnionType'\] | Type, Or |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='WildcardType'\] | Type, Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='Modifier'\]/self::\*\[@Token='public'\] | Visibility, World |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='Modifier'\]/self::\*\[@Token='protected'\] | Visibility, Subtype |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='Modifier'\]/self::\*\[@Token='private'\] | Visibility, Instance |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='Modifier'\]/self::\*\[@Token='abstract'\] | Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='Modifier'\]/self::\*\[@Token='static'\] | Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='Modifier'\]/self::\*\[@Token='final'\] | Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='Modifier'\]/self::\*\[@Token='strictfp'\] | Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='Modifier'\]/self::\*\[@Token='transient'\] | Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='Modifier'\]/self::\*\[@Token='volatile'\] | Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='Modifier'\]/self::\*\[@Token='synchronized'\] | Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='Modifier'\]/self::\*\[@Token='native'\] | Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='TryStatement'\] | Statement, Try |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='TryStatement'\]/\*\[@internalRole\]\[@internalRole='body'\] | Try, Body |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='TryStatement'\]/\*\[@internalRole\]\[@internalRole='catchClauses'\] | Try, Catch |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='TryStatement'\]/\*\[@internalRole\]\[@internalRole='finally'\] | Try, Finally |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='ThrowStatement'\] | Statement, Throw |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='AssertStatement'\] | Statement, Assert |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='MarkerAnnotation'\] | Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='MemberRef'\] | Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='MemberValuePair'\] | Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='MethodRef'\] | Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='MethodRefParameter'\] | Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='NormalAnnotation'\] | Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='SingleMemberAnnotation'\] | Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='TagElement'\] | Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='TextElement'\] | Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='BlockComment'\] | Comment |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='Javadoc'\] | Documentation, Comment |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='LineComment'\] | Comment |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='ArrayAccess'\] | Expression, Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='ArrayCreation'\] | Expression, Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='CastExpression'\] | Expression, Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='CreationReference'\] | Expression, Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='ExpressionMethodReference'\] | Expression, Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='ParenthesizedExpression'\] | Expression, Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='SuperMethodReference'\] | Expression, Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='ThisExpression'\] | Expression, This |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='Block'\] | Statement, Block, Scope |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='BreakStatement'\] | Statement, Break |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='EmptyStatement'\] | Statement |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='ExpressionStatement'\] | Statement |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='LabeledStatement'\] | Statement, Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='ReturnStatement'\] | Statement, Return |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='SynchronizedStatement'\] | Statement, Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='ArrayInitializer'\] | Expression, List, Literal |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='Dimension'\] | Type, Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='TypeParameter'\] | Type, Incomplete |
