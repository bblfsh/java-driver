| Path | Action |
|------|--------|
| /self::\*\[@InternalType='CompilationUnit'\] | File |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='QualifiedName'\] | QualifiedIdentifier, Expression |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='SimpleName'\] | SimpleIdentifier, Expression |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[\(@InternalType='MethodDeclaration'\) or \(@InternalType='TypeDeclaration'\)\]/self::\*\[child::\(@InternalType='Modifier'\) and \(@Token='public'\)\] | VisibleFromWorld |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[\(@InternalType='MethodDeclaration'\) or \(@InternalType='TypeDeclaration'\)\]/self::\*\[child::\(@InternalType='Modifier'\) and \(@Token='private'\)\] | VisibleFromType |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[\(@InternalType='MethodDeclaration'\) or \(@InternalType='TypeDeclaration'\)\]/self::\*\[child::\(@InternalType='Modifier'\) and \(@Token='protected'\)\] | VisibleFromSubtype |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[\(@InternalType='MethodDeclaration'\) or \(@InternalType='TypeDeclaration'\)\]/self::\*\[not\(child::\(@InternalType='Modifier'\) and \(\(@Token='public'\) or \(@Token='private'\) or \(@Token='protected'\)\)\)\] | VisibleFromPackage |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='PackageDeclaration'\] | PackageDeclaration |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='ImportDeclaration'\] | ImportDeclaration |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='ImportDeclaration'\]/\*\[@InternalType='QualifiedName'\] | ImportPath |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='AnonymousClassDeclaration'\] | TypeDeclaration, Expression, Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='AnonymousClassDeclaration'\]/\*\[@internalRole\]\[@internalRole='bodyDeclarations'\] | TypeDeclarationBody |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='AnnotationTypeDeclaration'\] | TypeDeclaration, Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='AnnotationTypeDeclaration'\]/\*\[@internalRole\]\[@internalRole='bodyDeclarations'\] | TypeDeclarationBody |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='EnumDeclaration'\] | TypeDeclaration, Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='TypeDeclaration'\] | TypeDeclaration |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='TypeDeclarationStatement'\] | TypeDeclaration, Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='MethodDeclaration'\] | FunctionDeclaration |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='MethodDeclaration'\]/\*\[@internalRole\]\[@internalRole='name'\] | FunctionDeclarationName |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='MethodDeclaration'\]/\*\[@internalRole\]\[@internalRole='body'\] | FunctionDeclarationBody |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='MethodDeclaration'\]/\*\[@internalRole\]\[@internalRole='parameters'\] | FunctionDeclarationArgument |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='MethodDeclaration'\]/\*\[@internalRole\]\[@internalRole='parameters'\]/self::\*\[@varargs\]\[@varargs='true'\] | FunctionDeclarationVarArgsList |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='MethodDeclaration'\]/\*\[@internalRole\]\[@internalRole='parameters'\]/\*\[@internalRole\]\[@internalRole='name'\] | FunctionDeclarationArgumentName |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='LambdaExpression'\] | FunctionDeclaration, Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='LambdaExpression'\]/\*\[@internalRole\]\[@internalRole='body'\] | FunctionDeclarationBody |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='LambdaExpression'\]/\*\[@internalRole\]\[@internalRole='parameters'\] | FunctionDeclarationArgument |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='LambdaExpression'\]/\*\[@internalRole\]\[@internalRole='parameters'\]/self::\*\[@varargs\]\[@varargs='true'\] | FunctionDeclarationVarArgsList |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='LambdaExpression'\]/\*\[@internalRole\]\[@internalRole='parameters'\]/\*\[@internalRole\]\[@internalRole='name'\] | FunctionDeclarationArgumentName |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='AnnotationTypeMemberDeclaration'\] | Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='EnumConstantDeclaration'\] | Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='FieldDeclaration'\] | Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='Initializer'\] | Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='SingleVariableDeclaration'\] | Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='VariableDeclarationExpression'\] | Expression, Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='VariableDeclarationFragment'\] | Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='VariableDeclarationStatement'\] | Statement, Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='BooleanLiteral'\] | BooleanLiteral, Expression |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='CharacterLiteral'\] | CharacterLiteral, Expression |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='NullLiteral'\] | NullLiteral, Expression |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='NumberLiteral'\] | NumberLiteral, Expression |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='StringLiteral'\] | StringLiteral, Expression |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='TypeLiteral'\] | TypeLiteral, Expression |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='ClassInstanceCreation'\] | Call, Expression, Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='ClassInstanceCreation'\]/\*\[@internalRole\]\[@internalRole='type'\] | CallCallee |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='ClassInstanceCreation'\]/\*\[@internalRole\]\[@internalRole='arguments'\] | CallPositionalArgument |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='ConstructorInvocation'\] | Call, Statement, Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='ConstructorInvocation'\]/\*\[@internalRole\]\[@internalRole='type'\] | CallCallee |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='ConstructorInvocation'\]/\*\[@internalRole\]\[@internalRole='arguments'\] | CallPositionalArgument |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='MethodInvocation'\] | Call, Expression |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='MethodInvocation'\]/\*\[@internalRole\]\[@internalRole='expression'\] | CallReceiver |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='MethodInvocation'\]/\*\[@internalRole\]\[@internalRole='name'\] | CallCallee |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='MethodInvocation'\]/\*\[@internalRole\]\[@internalRole='arguments'\] | CallPositionalArgument |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='SuperConstructorInvocation'\] | Call, Statement, Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='SuperConstructorInvocation'\]/\*\[@internalRole\]\[@internalRole='expression'\] | CallReceiver |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='SuperConstructorInvocation'\]/\*\[@internalRole\]\[@internalRole='arguments'\] | CallPositionalArgument |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='SuperMethodInvocation'\] | Call, Expression, Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='SuperMethodInvocation'\]/\*\[@internalRole\]\[@internalRole='qualifier'\] | CallCallee |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='SuperMethodInvocation'\]/\*\[@internalRole\]\[@internalRole='name'\] | CallCallee |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='SuperMethodInvocation'\]/\*\[@internalRole\]\[@internalRole='arguments'\] | CallPositionalArgument |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='IfStatement'\] | If, Statement |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='IfStatement'\]/\*\[@internalRole\]\[@internalRole='expression'\] | IfCondition |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='IfStatement'\]/\*\[@internalRole\]\[@internalRole='thenStatement'\] | IfBody |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='IfStatement'\]/\*\[@internalRole\]\[@internalRole='elseStatement'\] | IfElse |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='SwitchStatement'\] | Switch, Statement |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='SwitchStatement'\]/\*\[@InternalType='SwitchCase'\] | Statement |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='SwitchStatement'\]/\*\[@InternalType='SwitchCase'\]/self::\*\[child::\*\] | SwitchCase |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='SwitchStatement'\]/\*\[@InternalType='SwitchCase'\]/self::\*\[child::\*\]/\*\[@internalRole\]\[@internalRole='expression'\] | SwitchCaseCondition |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='SwitchStatement'\]/\*\[@InternalType='SwitchCase'\]/self::\*\[not\(child::\*\)\] | SwitchDefault |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='SwitchStatement'\]/\*\[@InternalType='ExpressionStatement'\] | SwitchCaseBody |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='EnhancedForStatement'\] | ForEach, Statement |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='EnhancedForStatement'\]/\*\[@internalRole\]\[@internalRole='parameter'\] | ForInit, ForUpdate |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='EnhancedForStatement'\]/\*\[@internalRole\]\[@internalRole='expression'\] | ForExpression |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='EnhancedForStatement'\]/\*\[@internalRole\]\[@internalRole='body'\] | ForBody |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='ForStatement'\] | For, Statement |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='ForStatement'\]/\*\[@internalRole\]\[@internalRole='initializers'\] | ForInit |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='ForStatement'\]/\*\[@internalRole\]\[@internalRole='expression'\] | ForExpression |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='ForStatement'\]/\*\[@internalRole\]\[@internalRole='updaters'\] | ForUpdate |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='ForStatement'\]/\*\[@internalRole\]\[@internalRole='body'\] | ForBody |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='WhileStatement'\] | While, Statement |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='WhileStatement'\]/\*\[@internalRole\]\[@internalRole='expression'\] | WhileCondition |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='WhileStatement'\]/\*\[@internalRole\]\[@internalRole='body'\] | WhileBody |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='DoStatement'\] | DoWhile, Statement |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='DoStatement'\]/\*\[@internalRole\]\[@internalRole='expression'\] | DoWhileCondition |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='DoStatement'\]/\*\[@internalRole\]\[@internalRole='body'\] | DoWhileBody |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='InfixExpression'\] | BinaryExpression, BinaryExpressionOp, Expression |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='InfixExpression'\]/self::\*\[@operator\]\[@operator='\+'\] | OpAdd |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='InfixExpression'\]/self::\*\[@operator\]\[@operator='\-'\] | OpSubstract |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='InfixExpression'\]/self::\*\[@operator\]\[@operator='\*'\] | OpMultiply |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='InfixExpression'\]/self::\*\[@operator\]\[@operator='/'\] | OpDivide |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='InfixExpression'\]/self::\*\[@operator\]\[@operator='%'\] | OpMod |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='InfixExpression'\]/self::\*\[@operator\]\[@operator='<<'\] | OpBitwiseLeftShift |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='InfixExpression'\]/self::\*\[@operator\]\[@operator='>>'\] | OpBitwiseRightShift |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='InfixExpression'\]/self::\*\[@operator\]\[@operator='>>>'\] | OpBitwiseUnsignedRightShift |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='InfixExpression'\]/self::\*\[@operator\]\[@operator='&'\] | OpBitwiseAnd |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='InfixExpression'\]/self::\*\[@operator\]\[@operator='\|'\] | OpBitwiseOr |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='InfixExpression'\]/self::\*\[@operator\]\[@operator='&&'\] | OpBooleanAnd |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='InfixExpression'\]/self::\*\[@operator\]\[@operator='\|\|'\] | OpBooleanOr |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='InfixExpression'\]/self::\*\[@operator\]\[@operator='^'\] | OpBooleanXor |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='InfixExpression'\]/\*\[@internalRole\]\[@internalRole='leftOperand'\] | BinaryExpressionLeft |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='InfixExpression'\]/\*\[@internalRole\]\[@internalRole='rightOperand'\] | BinaryExpressionRight |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='PostfixExpression'\] | Expression |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='PostfixExpression'\]/self::\*\[@operator\]\[@operator='\+\+'\] | OpPostIncrement |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='PostfixExpression'\]/self::\*\[@operator\]\[@operator='\-\-'\] | OpPostDecrement |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='PrefixExpression'\] | Expression |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='PrefixExpression'\]/self::\*\[@operator\]\[@operator='\+\+'\] | OpPreIncrement |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='PrefixExpression'\]/self::\*\[@operator\]\[@operator='\-\-'\] | OpPreDecrement |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='PrefixExpression'\]/self::\*\[@operator\]\[@operator='\+'\] | OpPositive |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='PrefixExpression'\]/self::\*\[@operator\]\[@operator='\-'\] | OpNegative |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='PrefixExpression'\]/self::\*\[@operator\]\[@operator='~'\] | OpBitwiseComplement |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='PrefixExpression'\]/self::\*\[@operator\]\[@operator='\!'\] | OpBooleanNot |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='Assignment'\] | Assignment, Expression |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='Assignment'\]/\*\[@internalRole\]\[@internalRole='leftHandSide'\] | AssignmentVariable |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='Assignment'\]/\*\[@internalRole\]\[@internalRole='rightHandSide'\] | AssignmentValue |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='Assignment'\]/self::\*\[not\(@operator\]\[@operator='='\)\] | AugmentedAssignmentOperator, AugmentedAssignment |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='Assignment'\]/self::\*\[not\(@operator\]\[@operator='='\)\]/self::\*\[@operator\]\[@operator='\+='\] | OpAdd |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='Assignment'\]/self::\*\[not\(@operator\]\[@operator='='\)\]/self::\*\[@operator\]\[@operator='\-='\] | OpSubstract |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='Assignment'\]/self::\*\[not\(@operator\]\[@operator='='\)\]/self::\*\[@operator\]\[@operator='\*='\] | OpMultiply |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='Assignment'\]/self::\*\[not\(@operator\]\[@operator='='\)\]/self::\*\[@operator\]\[@operator='/='\] | OpDivide |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='Assignment'\]/self::\*\[not\(@operator\]\[@operator='='\)\]/self::\*\[@operator\]\[@operator='%='\] | OpMod |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='Assignment'\]/self::\*\[not\(@operator\]\[@operator='='\)\]/self::\*\[@operator\]\[@operator='&='\] | OpBitwiseAnd |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='Assignment'\]/self::\*\[not\(@operator\]\[@operator='='\)\]/self::\*\[@operator\]\[@operator='\|='\] | OpBitwiseOr |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='Assignment'\]/self::\*\[not\(@operator\]\[@operator='='\)\]/self::\*\[@operator\]\[@operator='^='\] | OpBooleanXor |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='Assignment'\]/self::\*\[not\(@operator\]\[@operator='='\)\]/self::\*\[@operator\]\[@operator='<<='\] | OpBitwiseLeftShift |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='Assignment'\]/self::\*\[not\(@operator\]\[@operator='='\)\]/self::\*\[@operator\]\[@operator='>>='\] | OpBitwiseRightShift |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='Assignment'\]/self::\*\[not\(@operator\]\[@operator='='\)\]/self::\*\[@operator\]\[@operator='>>>='\] | OpBitwiseUnsignedRightShift |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='ArrayType'\] | Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='IntersectionType'\] | Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='NameQualifiedType'\] | Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='ParameterizedType'\] | Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='PrimitiveType'\] | Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='QualifiedType'\] | Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='SimpleType'\] | Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='UnionType'\] | Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='WildcardType'\] | Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='Modifier'\]/self::\*\[@Token='public'\] | VisibleFromWorld |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='Modifier'\]/self::\*\[@Token='protected'\] | VisibleFromSubtype |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='Modifier'\]/self::\*\[@Token='private'\] | VisibleFromInstance |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='Modifier'\]/self::\*\[@Token='abstract'\] | Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='Modifier'\]/self::\*\[@Token='static'\] | Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='Modifier'\]/self::\*\[@Token='final'\] | Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='Modifier'\]/self::\*\[@Token='strictfp'\] | Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='Modifier'\]/self::\*\[@Token='transient'\] | Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='Modifier'\]/self::\*\[@Token='volatile'\] | Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='Modifier'\]/self::\*\[@Token='synchronized'\] | Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='Modifier'\]/self::\*\[@Token='native'\] | Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='TryStatement'\] | Try, Statement |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='TryStatement'\]/\*\[@internalRole\]\[@internalRole='body'\] | TryBody |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='TryStatement'\]/\*\[@internalRole\]\[@internalRole='catchClauses'\] | TryCatch |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='TryStatement'\]/\*\[@internalRole\]\[@internalRole='finally'\] | TryFinally |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='ThrowStatement'\] | Throw, Statement |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='AssertStatement'\] | Assert, Statement |
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
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='ThisExpression'\] | This, Expression |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='Block'\] | BlockScope, Block, Statement |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='BreakStatement'\] | Break, Statement |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='EmptyStatement'\] | Statement |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='ExpressionStatement'\] | Statement |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='LabeledStatement'\] | Statement, Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='ReturnStatement'\] | Return, Statement |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='SynchronizedStatement'\] | Statement, Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='ArrayInitializer'\] | Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='Dimension'\] | Incomplete |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='TypeParameter'\] | Incomplete |
