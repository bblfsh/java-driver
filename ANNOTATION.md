| Path | Action |
|------|--------|
| /self::\*\[not\(@InternalType='CompilationUnit'\)\] | Error |
| /self::\*\[@InternalType='CompilationUnit'\] | File |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='QualifiedName'\] | QualifiedIdentifier |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='SimpleName'\] | SimpleIdentifier |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[\(@InternalType='MethodDeclaration'\) or \(@InternalType='TypeDeclaration'\)\]/self::\*\[child::\(@InternalType='Modifier'\) and \(@Token='public'\)\] | VisibleFromWorld |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[\(@InternalType='MethodDeclaration'\) or \(@InternalType='TypeDeclaration'\)\]/self::\*\[child::\(@InternalType='Modifier'\) and \(@Token='private'\)\] | VisibleFromType |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[\(@InternalType='MethodDeclaration'\) or \(@InternalType='TypeDeclaration'\)\]/self::\*\[child::\(@InternalType='Modifier'\) and \(@Token='protected'\)\] | VisibleFromSubtype |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[\(@InternalType='MethodDeclaration'\) or \(@InternalType='TypeDeclaration'\)\]/self::\*\[not\(child::\(@InternalType='Modifier'\) and \(\(@Token='public'\) or \(@Token='private'\) or \(@Token='protected'\)\)\)\] | VisibleFromPackage |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='PackageDeclaration'\] | PackageDeclaration |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='ImportDeclaration'\] | ImportDeclaration |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='ImportDeclaration'\]/\*\[@InternalType='QualifiedName'\] | ImportPath |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='TypeDeclaration'\] | TypeDeclaration |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='MethodDeclaration'\] | FunctionDeclaration |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='MethodDeclaration'\]/\*\[@internalRole\]\[@internalRole='name'\] | FunctionDeclarationName |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='MethodDeclaration'\]/\*\[@internalRole\]\[@internalRole='body'\] | FunctionDeclarationBody |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='MethodDeclaration'\]/\*\[@internalRole\]\[@internalRole='parameters'\] | FunctionDeclarationArgument |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='MethodDeclaration'\]/\*\[@internalRole\]\[@internalRole='parameters'\]/self::\*\[@varargs\]\[@varargs='true'\] | FunctionDeclarationVarArgsList |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='MethodDeclaration'\]/\*\[@internalRole\]\[@internalRole='parameters'\]/\*\[@internalRole\]\[@internalRole='name'\] | FunctionDeclarationArgumentName |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='BooleanLiteral'\] | BooleanLiteral |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='CharacterLiteral'\] | CharacterLiteral |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='NullLiteral'\] | NullLiteral |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='NumberLiteral'\] | NumberLiteral |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='StringLiteral'\] | StringLiteral |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='TypeLiteral'\] | TypeLiteral |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='MethodInvocation'\] | Call |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='MethodInvocation'\]/\*\[@internalRole\]\[@internalRole='expression'\] | CallReceiver |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='MethodInvocation'\]/\*\[@internalRole\]\[@internalRole='name'\] | CallCallee |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='MethodInvocation'\]/\*\[@internalRole\]\[@internalRole='arguments'\] | CallPositionalArgument |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='IfStatement'\] | If, Statement |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='IfStatement'\]/\*\[@internalRole\]\[@internalRole='expression'\] | IfCondition |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='IfStatement'\]/\*\[@internalRole\]\[@internalRole='thenStatement'\] | IfBody |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='IfStatement'\]/\*\[@internalRole\]\[@internalRole='elseStatement'\] | IfElse |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='SwitchStatement'\] | Switch, Statement |
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
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='InfixExpression'\] | BinaryExpression, BinaryExpressionOp |
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
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='PostfixExpression'\]/self::\*\[@operator\]\[@operator='\+\+'\] | OpPostIncrement |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='PostfixExpression'\]/self::\*\[@operator\]\[@operator='\-\-'\] | OpPostDecrement |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='PrefixExpression'\]/self::\*\[@operator\]\[@operator='\+\+'\] | OpPreIncrement |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='PrefixExpression'\]/self::\*\[@operator\]\[@operator='\-\-'\] | OpPreDecrement |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='PrefixExpression'\]/self::\*\[@operator\]\[@operator='\+'\] | OpPositive |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='PrefixExpression'\]/self::\*\[@operator\]\[@operator='\-'\] | OpNegative |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='PrefixExpression'\]/self::\*\[@operator\]\[@operator='~'\] | OpBitwiseComplement |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='PrefixExpression'\]/self::\*\[@operator\]\[@operator='\!'\] | OpBooleanNot |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='Assignment'\] | Assignment |
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
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='TryStatement'\] | Try, Statement |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='TryStatement'\]/\*\[@internalRole\]\[@internalRole='body'\] | TryBody |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='TryStatement'\]/\*\[@internalRole\]\[@internalRole='catchClauses'\] | TryCatch |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='TryStatement'\]/\*\[@internalRole\]\[@internalRole='finally'\] | TryFinally |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='ThrowStatement'\] | Throw, Statement |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='AssertStatement'\] | Assert, Statement |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='Block'\] | BlockScope, Block |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='ExpressionStatement'\] | Statement |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='ReturnStatement'\] | Return, Statement |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='BreakStatement'\] | Break, Statement |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='ThisExpression'\] | This, Expression |
| /self::\*\[@InternalType='CompilationUnit'\]//\*\[@InternalType='Javadoc'\] | Documentation, Comment |
