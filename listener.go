// generated by Textmapper; DO NOT EDIT

package ll

import (
	"fmt"
)

type NodeType int

type Listener func(t NodeType, offset, endoffset int)

const (
	NoType NodeType = iota
	GlobalIdent
	LocalIdent
	LabelIdent
	AttrGroupID
	ComdatName
	MetadataName
	MetadataID
	BoolLit
	IntLit
	UintLit
	FloatLit
	StringLit
	NullLit
	Module           // TargetDefs=(TargetDef)* TopLevelEntities=(TopLevelEntity)*
	SourceFilename   // Name=StringLit
	TargetDataLayout // DataLayout=StringLit
	TargetTriple     // TargetTriple=StringLit
	ModuleAsm        // Asm=StringLit
	TypeDef          // Name=LocalIdent Typ=(OpaqueType | Type)
	ComdatDef        // Name=ComdatName Kind=SelectionKind
	SelectionKind
	GlobalDecl // Name=GlobalIdent Linkage=(ExternLinkage | Linkage)? Preemption? Visibility? DLLStorageClass? ThreadLocal? UnnamedAddr? AddrSpace? ExternallyInitialized? Immutable ContentType=Type Init=Constant? GlobalFields=(GlobalField)* Metadata=(MetadataAttachment)* FuncAttrs=(FuncAttribute)*
	ExternallyInitialized
	Immutable
	IndirectSymbolDef // Name=GlobalIdent ExternLinkage? Linkage? Preemption? Visibility? DLLStorageClass? ThreadLocal? UnnamedAddr? IndirectSymbolKind ContentType=Type IndirectSymbol Partitions=(Partition)*
	IndirectSymbolKind
	FuncDecl         // Metadata=(MetadataAttachment)* Header=FuncHeader
	FuncDef          // Header=FuncHeader Metadata=(MetadataAttachment)* Body=FuncBody
	FuncHeader       // ExternLinkage? Linkage? Preemption? Visibility? DLLStorageClass? CallingConv? ReturnAttrs=(ReturnAttribute)* RetType=Type Name=GlobalIdent Params UnnamedAddr? AddrSpace? FuncHdrFields=(FuncHdrField)*
	GCNode           // Name=StringLit
	Prefix           // TypeConst
	Prologue         // TypeConst
	Personality      // TypeConst
	FuncBody         // Blocks=(BasicBlock)+ UseListOrders=(UseListOrder)*
	AttrGroupDef     // ID=AttrGroupID FuncAttrs=(FuncAttribute)*
	NamedMetadataDef // Name=MetadataName MDNodes=(MetadataNode)*
	MetadataDef      // ID=MetadataID Distinct? MDNode=(MDTuple | SpecializedMDNode)
	Distinct
	UseListOrder   // Val=TypeValue Indices=(UintLit)+
	UseListOrderBB // Func=GlobalIdent Block=LocalIdent Indices=(UintLit)+
	VoidType
	FuncType // RetType=Type Params
	IntType
	FloatType // FloatKind
	FloatKind
	MMXType
	PointerType        // Elem=Type AddrSpace?
	VectorType         // Len=UintLit Elem=Type
	ScalableVectorType // Len=UintLit Elem=Type
	LabelType
	TokenType
	MetadataType
	ArrayType        // Len=UintLit Elem=Type
	StructType       // Fields=(Type)*
	PackedStructType // Fields=(Type)*
	OpaqueType
	NamedType // Name=LocalIdent
	InlineAsm // SideEffect? AlignStackTok? IntelDialect? Asm=StringLit Constraints=StringLit
	SideEffect
	AlignStackTok
	IntelDialect
	BoolConst  // BoolLit
	IntConst   // IntLit
	FloatConst // FloatLit
	NullConst  // NullLit
	NoneConst
	StructConst    // Fields=(TypeConst)*
	ArrayConst     // Elems=(TypeConst)*
	CharArrayConst // Val=StringLit
	VectorConst    // Elems=(TypeConst)*
	ZeroInitializerConst
	UndefConst
	BlockAddressConst  // Func=GlobalIdent Block=LocalIdent
	FNegExpr           // X=TypeConst
	AddExpr            // OverflowFlags=(OverflowFlag)* X=TypeConst Y=TypeConst
	FAddExpr           // X=TypeConst Y=TypeConst
	SubExpr            // OverflowFlags=(OverflowFlag)* X=TypeConst Y=TypeConst
	FSubExpr           // X=TypeConst Y=TypeConst
	MulExpr            // OverflowFlags=(OverflowFlag)* X=TypeConst Y=TypeConst
	FMulExpr           // X=TypeConst Y=TypeConst
	UDivExpr           // Exact? X=TypeConst Y=TypeConst
	SDivExpr           // Exact? X=TypeConst Y=TypeConst
	FDivExpr           // X=TypeConst Y=TypeConst
	URemExpr           // X=TypeConst Y=TypeConst
	SRemExpr           // X=TypeConst Y=TypeConst
	FRemExpr           // X=TypeConst Y=TypeConst
	ShlExpr            // OverflowFlags=(OverflowFlag)* X=TypeConst Y=TypeConst
	LShrExpr           // Exact? X=TypeConst Y=TypeConst
	AShrExpr           // Exact? X=TypeConst Y=TypeConst
	AndExpr            // X=TypeConst Y=TypeConst
	OrExpr             // X=TypeConst Y=TypeConst
	XorExpr            // X=TypeConst Y=TypeConst
	ExtractElementExpr // X=TypeConst Index=TypeConst
	InsertElementExpr  // X=TypeConst Elem=TypeConst Index=TypeConst
	ShuffleVectorExpr  // X=TypeConst Y=TypeConst Mask=TypeConst
	ExtractValueExpr   // X=TypeConst Indices=(UintLit)*
	InsertValueExpr    // X=TypeConst Elem=TypeConst Indices=(UintLit)*
	GetElementPtrExpr  // InBounds? ElemType=Type Src=TypeConst Indices=(GEPIndex)*
	GEPIndex           // InRange? Index=TypeConst
	InRange
	TruncExpr          // From=TypeConst To=Type
	ZExtExpr           // From=TypeConst To=Type
	SExtExpr           // From=TypeConst To=Type
	FPTruncExpr        // From=TypeConst To=Type
	FPExtExpr          // From=TypeConst To=Type
	FPToUIExpr         // From=TypeConst To=Type
	FPToSIExpr         // From=TypeConst To=Type
	UIToFPExpr         // From=TypeConst To=Type
	SIToFPExpr         // From=TypeConst To=Type
	PtrToIntExpr       // From=TypeConst To=Type
	IntToPtrExpr       // From=TypeConst To=Type
	BitCastExpr        // From=TypeConst To=Type
	AddrSpaceCastExpr  // From=TypeConst To=Type
	ICmpExpr           // Pred=IPred X=TypeConst Y=TypeConst
	FCmpExpr           // Pred=FPred X=TypeConst Y=TypeConst
	SelectExpr         // Cond=TypeConst X=TypeConst Y=TypeConst
	BasicBlock         // Name=LabelIdent? Insts=(Instruction)* Term=Terminator
	LocalDefInst       // Name=LocalIdent Inst=ValueInstruction
	FNegInst           // FastMathFlags=(FastMathFlag)* X=TypeValue Metadata=(MetadataAttachment)*
	AddInst            // OverflowFlags=(OverflowFlag)* X=TypeValue Y=Value Metadata=(MetadataAttachment)*
	FAddInst           // FastMathFlags=(FastMathFlag)* X=TypeValue Y=Value Metadata=(MetadataAttachment)*
	SubInst            // OverflowFlags=(OverflowFlag)* X=TypeValue Y=Value Metadata=(MetadataAttachment)*
	FSubInst           // FastMathFlags=(FastMathFlag)* X=TypeValue Y=Value Metadata=(MetadataAttachment)*
	MulInst            // OverflowFlags=(OverflowFlag)* X=TypeValue Y=Value Metadata=(MetadataAttachment)*
	FMulInst           // FastMathFlags=(FastMathFlag)* X=TypeValue Y=Value Metadata=(MetadataAttachment)*
	UDivInst           // Exact? X=TypeValue Y=Value Metadata=(MetadataAttachment)*
	SDivInst           // Exact? X=TypeValue Y=Value Metadata=(MetadataAttachment)*
	FDivInst           // FastMathFlags=(FastMathFlag)* X=TypeValue Y=Value Metadata=(MetadataAttachment)*
	URemInst           // X=TypeValue Y=Value Metadata=(MetadataAttachment)*
	SRemInst           // X=TypeValue Y=Value Metadata=(MetadataAttachment)*
	FRemInst           // FastMathFlags=(FastMathFlag)* X=TypeValue Y=Value Metadata=(MetadataAttachment)*
	ShlInst            // OverflowFlags=(OverflowFlag)* X=TypeValue Y=Value Metadata=(MetadataAttachment)*
	LShrInst           // Exact? X=TypeValue Y=Value Metadata=(MetadataAttachment)*
	AShrInst           // Exact? X=TypeValue Y=Value Metadata=(MetadataAttachment)*
	AndInst            // X=TypeValue Y=Value Metadata=(MetadataAttachment)*
	OrInst             // X=TypeValue Y=Value Metadata=(MetadataAttachment)*
	XorInst            // X=TypeValue Y=Value Metadata=(MetadataAttachment)*
	ExtractElementInst // X=TypeValue Index=TypeValue Metadata=(MetadataAttachment)*
	InsertElementInst  // X=TypeValue Elem=TypeValue Index=TypeValue Metadata=(MetadataAttachment)*
	ShuffleVectorInst  // X=TypeValue Y=TypeValue Mask=TypeValue Metadata=(MetadataAttachment)*
	ExtractValueInst   // X=TypeValue Indices=(UintLit)+ Metadata=(MetadataAttachment)*
	InsertValueInst    // X=TypeValue Elem=TypeValue Indices=(UintLit)+ Metadata=(MetadataAttachment)*
	AllocaInst         // InAlloca? SwiftError? ElemType=Type NElems=TypeValue? Align? AddrSpace? Metadata=(MetadataAttachment)*
	InAlloca
	SwiftError
	LoadInst    // Atomic? Volatile? ElemType=Type Src=TypeValue SyncScope? Ordering=AtomicOrdering? Align? Metadata=(MetadataAttachment)*
	StoreInst   // Atomic? Volatile? Src=TypeValue Dst=TypeValue SyncScope? Ordering=AtomicOrdering? Align? Metadata=(MetadataAttachment)*
	FenceInst   // SyncScope? Ordering=AtomicOrdering Metadata=(MetadataAttachment)*
	CmpXchgInst // Weak? Volatile? Ptr=TypeValue Cmp=TypeValue New=TypeValue SyncScope? SuccessOrdering=AtomicOrdering FailureOrdering=AtomicOrdering Metadata=(MetadataAttachment)*
	Weak
	AtomicRMWInst // Volatile? Op=AtomicOp Dst=TypeValue X=TypeValue SyncScope? Ordering=AtomicOrdering Metadata=(MetadataAttachment)*
	AtomicOp
	GetElementPtrInst // InBounds? ElemType=Type Src=TypeValue Indices=(TypeValue)* Metadata=(MetadataAttachment)*
	TruncInst         // From=TypeValue To=Type Metadata=(MetadataAttachment)*
	ZExtInst          // From=TypeValue To=Type Metadata=(MetadataAttachment)*
	SExtInst          // From=TypeValue To=Type Metadata=(MetadataAttachment)*
	FPTruncInst       // From=TypeValue To=Type Metadata=(MetadataAttachment)*
	FPExtInst         // From=TypeValue To=Type Metadata=(MetadataAttachment)*
	FPToUIInst        // From=TypeValue To=Type Metadata=(MetadataAttachment)*
	FPToSIInst        // From=TypeValue To=Type Metadata=(MetadataAttachment)*
	UIToFPInst        // From=TypeValue To=Type Metadata=(MetadataAttachment)*
	SIToFPInst        // From=TypeValue To=Type Metadata=(MetadataAttachment)*
	PtrToIntInst      // From=TypeValue To=Type Metadata=(MetadataAttachment)*
	IntToPtrInst      // From=TypeValue To=Type Metadata=(MetadataAttachment)*
	BitCastInst       // From=TypeValue To=Type Metadata=(MetadataAttachment)*
	AddrSpaceCastInst // From=TypeValue To=Type Metadata=(MetadataAttachment)*
	ICmpInst          // Pred=IPred X=TypeValue Y=Value Metadata=(MetadataAttachment)*
	FCmpInst          // FastMathFlags=(FastMathFlag)* Pred=FPred X=TypeValue Y=Value Metadata=(MetadataAttachment)*
	PhiInst           // FastMathFlags=(FastMathFlag)* Typ=Type Incs=(Inc)+ Metadata=(MetadataAttachment)*
	Inc               // X=Value Pred=LocalIdent
	SelectInst        // FastMathFlags=(FastMathFlag)* Cond=TypeValue ValueTrue=TypeValue ValueFalse=TypeValue Metadata=(MetadataAttachment)*
	FreezeInst        // X=TypeValue
	CallInst          // Tail? FastMathFlags=(FastMathFlag)* CallingConv? ReturnAttrs=(ReturnAttribute)* AddrSpace? Typ=Type Callee=Value Args FuncAttrs=(FuncAttribute)* OperandBundles=(OperandBundle)* Metadata=(MetadataAttachment)*
	Tail
	VAArgInst      // ArgList=TypeValue ArgType=Type Metadata=(MetadataAttachment)*
	LandingPadInst // ResultType=Type Cleanup? Clauses=(Clause)* Metadata=(MetadataAttachment)*
	Cleanup
	Clause // ClauseType X=TypeValue
	ClauseType
	CatchPadInst               // CatchSwitch=LocalIdent Args=(ExceptionArg)* Metadata=(MetadataAttachment)*
	CleanupPadInst             // ParentPad=ExceptionPad Args=(ExceptionArg)* Metadata=(MetadataAttachment)*
	LocalDefTerm               // Name=LocalIdent Term=ValueTerminator
	RetTerm                    // XTyp=(ConcreteType | VoidType) X=Value? Metadata=(MetadataAttachment)*
	BrTerm                     // Target=Label Metadata=(MetadataAttachment)*
	CondBrTerm                 // CondTyp=IntType Cond=Value TargetTrue=Label TargetFalse=Label Metadata=(MetadataAttachment)*
	SwitchTerm                 // X=TypeValue Default=Label Cases=(Case)* Metadata=(MetadataAttachment)*
	Case                       // X=TypeConst Target=Label
	IndirectBrTerm             // Addr=TypeValue ValidTargets=(Label)* Metadata=(MetadataAttachment)*
	InvokeTerm                 // CallingConv? ReturnAttrs=(ReturnAttribute)* AddrSpace? Typ=Type Invokee=Value Args FuncAttrs=(FuncAttribute)* OperandBundles=(OperandBundle)* NormalRetTarget=Label ExceptionRetTarget=Label Metadata=(MetadataAttachment)*
	CallBrTerm                 // CallingConv? ReturnAttrs=(ReturnAttribute)* AddrSpace? Typ=Type Callee=Value Args FuncAttrs=(FuncAttribute)* OperandBundles=(OperandBundle)* NormalRetTarget=Label OtherRetTargets=(Label)* Metadata=(MetadataAttachment)*
	ResumeTerm                 // X=TypeValue Metadata=(MetadataAttachment)*
	CatchSwitchTerm            // ParentPad=ExceptionPad Handlers DefaultUnwindTarget=UnwindTarget Metadata=(MetadataAttachment)*
	Handlers                   // Labels=(Label)+
	CatchRetTerm               // CatchPad=Value Target=Label Metadata=(MetadataAttachment)*
	CleanupRetTerm             // CleanupPad=Value UnwindTarget Metadata=(MetadataAttachment)*
	UnreachableTerm            // Metadata=(MetadataAttachment)*
	MDTuple                    // MDFields=(MDField)*
	MDString                   // Val=StringLit
	MetadataAttachment         // Name=MetadataName MDNode
	DIBasicType                // Fields=(DIBasicTypeField)*
	DICommonBlock              // Fields=(DICommonBlockField)*
	DICompileUnit              // Fields=(DICompileUnitField)*
	DICompositeType            // Fields=(DICompositeTypeField)*
	DIDerivedType              // Fields=(DIDerivedTypeField)*
	DIEnumerator               // Fields=(DIEnumeratorField)*
	DIExpression               // Fields=(DIExpressionField)*
	DIFile                     // Fields=(DIFileField)*
	DIGlobalVariable           // Fields=(DIGlobalVariableField)*
	DIGlobalVariableExpression // Fields=(DIGlobalVariableExpressionField)*
	DIImportedEntity           // Fields=(DIImportedEntityField)*
	DILabel                    // Fields=(DILabelField)*
	DILexicalBlock             // Fields=(DILexicalBlockField)*
	DILexicalBlockFile         // Fields=(DILexicalBlockFileField)*
	DILocalVariable            // Fields=(DILocalVariableField)*
	DILocation                 // Fields=(DILocationField)*
	DIMacro                    // Fields=(DIMacroField)*
	DIMacroFile                // Fields=(DIMacroFileField)*
	DIModule                   // Fields=(DIModuleField)*
	DINamespace                // Fields=(DINamespaceField)*
	DIObjCProperty             // Fields=(DIObjCPropertyField)*
	DISubprogram               // Fields=(DISubprogramField)*
	DISubrange                 // Fields=(DISubrangeField)*
	DISubroutineType           // Fields=(DISubroutineTypeField)*
	DITemplateTypeParameter    // Fields=(DITemplateTypeParameterField)*
	DITemplateValueParameter   // Fields=(DITemplateValueParameterField)*
	GenericDINode              // Fields=(GenericDINodeField)*
	AlignField                 // Align=UintLit
	ArgField                   // Arg=UintLit
	AttributesField            // Attributes=UintLit
	BaseTypeField              // BaseType=MDField
	CCField                    // CC=DwarfCC
	ChecksumField              // Checksum=StringLit
	ChecksumkindField          // Checksumkind=ChecksumKind
	ColumnField                // Column=IntLit
	ConfigMacrosField          // ConfigMacros=StringLit
	ContainingTypeField        // ContainingType=MDField
	CountField                 // Count=MDFieldOrInt
	DebugInfoForProfilingField // DebugInfoForProfiling=BoolLit
	DeclarationField           // Declaration=MDField
	DirectoryField             // Directory=StringLit
	DiscriminatorField         // Discriminator=MDField
	DataLocationField          // DataLocation=MDField
	DefaultedField             // Name=BoolLit
	DiscriminatorIntField      // Discriminator=UintLit
	DwarfAddressSpaceField     // DwarfAddressSpace=UintLit
	DwoIdField                 // DwoId=UintLit
	ElementsField              // Elements=MDField
	EmissionKindField          // EmissionKind
	EncodingField              // Encoding=DwarfAttEncodingOrUint
	EntityField                // Entity=MDField
	EnumsField                 // Enums=MDField
	ExportSymbolsField         // ExportSymbols=BoolLit
	ExprField                  // Expr=MDField
	ExtraDataField             // ExtraData=MDField
	FileField                  // File=MDField
	FilenameField              // Filename=StringLit
	FlagsField                 // Flags=DIFlags
	FlagsStringField           // Flags=StringLit
	GetterField                // Getter=StringLit
	GlobalsField               // Globals=MDField
	HeaderField                // Header=StringLit
	IdentifierField            // Identifier=StringLit
	ImportsField               // Imports=MDField
	IncludePathField           // IncludePath=StringLit
	InlinedAtField             // InlinedAt=MDField
	IsDefinitionField          // IsDefinition=BoolLit
	IsImplicitCodeField        // IsImplicitCode=BoolLit
	IsLocalField               // IsLocal=BoolLit
	IsOptimizedField           // IsOptimized=BoolLit
	IsUnsignedField            // IsUnsigned=BoolLit
	APINotesField              // APINotes=StringLit
	LanguageField              // Language=DwarfLang
	LineField                  // Line=IntLit
	LinkageNameField           // LinkageName=StringLit
	LowerBoundField            // LowerBound=MDFieldOrInt
	MacrosField                // Macros=MDField
	NameField                  // Name=StringLit
	NameTableKindField         // NameTableKind
	NodesField                 // Nodes=MDField
	OffsetField                // OffsetField=UintLit
	OperandsField              // Operands=(MDField)*
	ProducerField              // Producer=StringLit
	RangesBaseAddressField     // RangesBaseAddress=BoolLit
	RetainedNodesField         // RetainedNodes=MDField
	RetainedTypesField         // RetainedTypes=MDField
	RuntimeLangField           // RuntimeLang=DwarfLang
	RuntimeVersionField        // RuntimeVersion=UintLit
	ScopeField                 // Scope=MDField
	ScopeLineField             // ScopeLine=IntLit
	SDKField                   // SDK=StringLit
	SetterField                // Setter=StringLit
	SizeField                  // Size=UintLit
	SourceField                // Source=StringLit
	SPFlagsField               // SPFlags=DISPFlags
	SplitDebugFilenameField    // SplitDebugFilename=StringLit
	SplitDebugInliningField    // SplitDebugInlining=BoolLit
	StrideField                // Stride=MDFieldOrInt
	SysrootField               // Sysroot=StringLit
	TagField                   // Tag=DwarfTag
	TemplateParamsField        // TemplateParams=MDField
	ThisAdjustmentField        // ThisAdjustment=IntLit
	ThrownTypesField           // ThrownTypes=MDField
	TypeField                  // Typ=MDField
	TypeMacinfoField           // Typ=DwarfMacinfo
	TypesField                 // Types=MDField
	UnitField                  // Unit=MDField
	UpperBoundField            // UpperBound=MDFieldOrInt
	ValueField                 // Value=MDField
	ValueIntField              // Value=IntLit
	ValueStringField           // Value=StringLit
	VarField                   // Var=MDField
	VirtualIndexField          // VirtualIndex=UintLit
	VirtualityField            // Virtuality=DwarfVirtuality
	VtableHolderField          // VtableHolder=MDField
	ChecksumKind
	DIFlags // Flags=(DIFlag)+
	DIFlagEnum
	DIFlagInt // UintLit
	DISPFlags // Flags=(DISPFlag)+
	DISPFlagEnum
	DISPFlagInt // UintLit
	DwarfAttEncodingEnum
	DwarfAttEncodingInt // UintLit
	DwarfCCEnum
	DwarfCCInt // UintLit
	DwarfLangEnum
	DwarfLangInt // UintLit
	DwarfMacinfoEnum
	DwarfMacinfoInt // UintLit
	DwarfOp
	DwarfTagEnum
	DwarfTagInt // UintLit
	DwarfVirtualityEnum
	DwarfVirtualityInt // UintLit
	EmissionKindEnum
	EmissionKindInt // UintLit
	NameTableKindEnum
	NameTableKindInt // UintLit
	AddrSpace        // N=UintLit
	Align            // N=UintLit
	AlignPair        // N=UintLit
	AlignStack       // N=UintLit
	AlignStackPair   // N=UintLit
	AllocSize        // ElemSizeIndex=UintLit NElemsIndex=UintLit?
	Args             // Args=(Arg)*
	Arg              // Typ=(ConcreteType | MetadataType) Attrs=(ParamAttribute)* Val=(Metadata | Value)
	Atomic
	AtomicOrdering
	AttrPair   // Key=StringLit Val=StringLit
	AttrString // Val=StringLit
	Byval      // Typ=Type?
	CallingConvEnum
	CallingConvInt        // UintLit
	Comdat                // Name=ComdatName?
	Dereferenceable       // N=UintLit
	DereferenceableOrNull // N=UintLit
	DLLStorageClass
	Ellipsis
	Exact
	ExceptionArg // Typ=(ConcreteType | MetadataType) Val=(Metadata | Value)
	FastMathFlag
	FPred
	FuncAttr
	InBounds
	IPred
	Label // Typ=LabelType Name=LocalIdent
	Linkage
	ExternLinkage
	OperandBundle // Tag=StringLit Inputs=(TypeValue)*
	OverflowFlag
	Params // Params=(Param)* Variadic=Ellipsis?
	Param  // Typ=Type Attrs=(ParamAttribute)* Name=LocalIdent?
	ParamAttr
	Partition    // Name=StringLit
	Preallocated // Typ=Type
	Preemption
	ReturnAttr
	Section     // Name=StringLit
	SyncScope   // Scope=StringLit
	ThreadLocal // Model=TLSModel?
	TLSModel
	TypeConst // Typ=FirstClassType Val=Constant
	TypeValue // Typ=FirstClassType Val=Value
	UnnamedAddr
	UnwindToCaller
	Visibility
	Volatile
	NodeTypeMax
)

var nodeTypeStr = [...]string{
	"NONE",
	"GlobalIdent",
	"LocalIdent",
	"LabelIdent",
	"AttrGroupID",
	"ComdatName",
	"MetadataName",
	"MetadataID",
	"BoolLit",
	"IntLit",
	"UintLit",
	"FloatLit",
	"StringLit",
	"NullLit",
	"Module",
	"SourceFilename",
	"TargetDataLayout",
	"TargetTriple",
	"ModuleAsm",
	"TypeDef",
	"ComdatDef",
	"SelectionKind",
	"GlobalDecl",
	"ExternallyInitialized",
	"Immutable",
	"IndirectSymbolDef",
	"IndirectSymbolKind",
	"FuncDecl",
	"FuncDef",
	"FuncHeader",
	"GCNode",
	"Prefix",
	"Prologue",
	"Personality",
	"FuncBody",
	"AttrGroupDef",
	"NamedMetadataDef",
	"MetadataDef",
	"Distinct",
	"UseListOrder",
	"UseListOrderBB",
	"VoidType",
	"FuncType",
	"IntType",
	"FloatType",
	"FloatKind",
	"MMXType",
	"PointerType",
	"VectorType",
	"ScalableVectorType",
	"LabelType",
	"TokenType",
	"MetadataType",
	"ArrayType",
	"StructType",
	"PackedStructType",
	"OpaqueType",
	"NamedType",
	"InlineAsm",
	"SideEffect",
	"AlignStackTok",
	"IntelDialect",
	"BoolConst",
	"IntConst",
	"FloatConst",
	"NullConst",
	"NoneConst",
	"StructConst",
	"ArrayConst",
	"CharArrayConst",
	"VectorConst",
	"ZeroInitializerConst",
	"UndefConst",
	"BlockAddressConst",
	"FNegExpr",
	"AddExpr",
	"FAddExpr",
	"SubExpr",
	"FSubExpr",
	"MulExpr",
	"FMulExpr",
	"UDivExpr",
	"SDivExpr",
	"FDivExpr",
	"URemExpr",
	"SRemExpr",
	"FRemExpr",
	"ShlExpr",
	"LShrExpr",
	"AShrExpr",
	"AndExpr",
	"OrExpr",
	"XorExpr",
	"ExtractElementExpr",
	"InsertElementExpr",
	"ShuffleVectorExpr",
	"ExtractValueExpr",
	"InsertValueExpr",
	"GetElementPtrExpr",
	"GEPIndex",
	"InRange",
	"TruncExpr",
	"ZExtExpr",
	"SExtExpr",
	"FPTruncExpr",
	"FPExtExpr",
	"FPToUIExpr",
	"FPToSIExpr",
	"UIToFPExpr",
	"SIToFPExpr",
	"PtrToIntExpr",
	"IntToPtrExpr",
	"BitCastExpr",
	"AddrSpaceCastExpr",
	"ICmpExpr",
	"FCmpExpr",
	"SelectExpr",
	"BasicBlock",
	"LocalDefInst",
	"FNegInst",
	"AddInst",
	"FAddInst",
	"SubInst",
	"FSubInst",
	"MulInst",
	"FMulInst",
	"UDivInst",
	"SDivInst",
	"FDivInst",
	"URemInst",
	"SRemInst",
	"FRemInst",
	"ShlInst",
	"LShrInst",
	"AShrInst",
	"AndInst",
	"OrInst",
	"XorInst",
	"ExtractElementInst",
	"InsertElementInst",
	"ShuffleVectorInst",
	"ExtractValueInst",
	"InsertValueInst",
	"AllocaInst",
	"InAlloca",
	"SwiftError",
	"LoadInst",
	"StoreInst",
	"FenceInst",
	"CmpXchgInst",
	"Weak",
	"AtomicRMWInst",
	"AtomicOp",
	"GetElementPtrInst",
	"TruncInst",
	"ZExtInst",
	"SExtInst",
	"FPTruncInst",
	"FPExtInst",
	"FPToUIInst",
	"FPToSIInst",
	"UIToFPInst",
	"SIToFPInst",
	"PtrToIntInst",
	"IntToPtrInst",
	"BitCastInst",
	"AddrSpaceCastInst",
	"ICmpInst",
	"FCmpInst",
	"PhiInst",
	"Inc",
	"SelectInst",
	"FreezeInst",
	"CallInst",
	"Tail",
	"VAArgInst",
	"LandingPadInst",
	"Cleanup",
	"Clause",
	"ClauseType",
	"CatchPadInst",
	"CleanupPadInst",
	"LocalDefTerm",
	"RetTerm",
	"BrTerm",
	"CondBrTerm",
	"SwitchTerm",
	"Case",
	"IndirectBrTerm",
	"InvokeTerm",
	"CallBrTerm",
	"ResumeTerm",
	"CatchSwitchTerm",
	"Handlers",
	"CatchRetTerm",
	"CleanupRetTerm",
	"UnreachableTerm",
	"MDTuple",
	"MDString",
	"MetadataAttachment",
	"DIBasicType",
	"DICommonBlock",
	"DICompileUnit",
	"DICompositeType",
	"DIDerivedType",
	"DIEnumerator",
	"DIExpression",
	"DIFile",
	"DIGlobalVariable",
	"DIGlobalVariableExpression",
	"DIImportedEntity",
	"DILabel",
	"DILexicalBlock",
	"DILexicalBlockFile",
	"DILocalVariable",
	"DILocation",
	"DIMacro",
	"DIMacroFile",
	"DIModule",
	"DINamespace",
	"DIObjCProperty",
	"DISubprogram",
	"DISubrange",
	"DISubroutineType",
	"DITemplateTypeParameter",
	"DITemplateValueParameter",
	"GenericDINode",
	"AlignField",
	"ArgField",
	"AttributesField",
	"BaseTypeField",
	"CCField",
	"ChecksumField",
	"ChecksumkindField",
	"ColumnField",
	"ConfigMacrosField",
	"ContainingTypeField",
	"CountField",
	"DebugInfoForProfilingField",
	"DeclarationField",
	"DirectoryField",
	"DiscriminatorField",
	"DataLocationField",
	"DefaultedField",
	"DiscriminatorIntField",
	"DwarfAddressSpaceField",
	"DwoIdField",
	"ElementsField",
	"EmissionKindField",
	"EncodingField",
	"EntityField",
	"EnumsField",
	"ExportSymbolsField",
	"ExprField",
	"ExtraDataField",
	"FileField",
	"FilenameField",
	"FlagsField",
	"FlagsStringField",
	"GetterField",
	"GlobalsField",
	"HeaderField",
	"IdentifierField",
	"ImportsField",
	"IncludePathField",
	"InlinedAtField",
	"IsDefinitionField",
	"IsImplicitCodeField",
	"IsLocalField",
	"IsOptimizedField",
	"IsUnsignedField",
	"APINotesField",
	"LanguageField",
	"LineField",
	"LinkageNameField",
	"LowerBoundField",
	"MacrosField",
	"NameField",
	"NameTableKindField",
	"NodesField",
	"OffsetField",
	"OperandsField",
	"ProducerField",
	"RangesBaseAddressField",
	"RetainedNodesField",
	"RetainedTypesField",
	"RuntimeLangField",
	"RuntimeVersionField",
	"ScopeField",
	"ScopeLineField",
	"SDKField",
	"SetterField",
	"SizeField",
	"SourceField",
	"SPFlagsField",
	"SplitDebugFilenameField",
	"SplitDebugInliningField",
	"StrideField",
	"SysrootField",
	"TagField",
	"TemplateParamsField",
	"ThisAdjustmentField",
	"ThrownTypesField",
	"TypeField",
	"TypeMacinfoField",
	"TypesField",
	"UnitField",
	"UpperBoundField",
	"ValueField",
	"ValueIntField",
	"ValueStringField",
	"VarField",
	"VirtualIndexField",
	"VirtualityField",
	"VtableHolderField",
	"ChecksumKind",
	"DIFlags",
	"DIFlagEnum",
	"DIFlagInt",
	"DISPFlags",
	"DISPFlagEnum",
	"DISPFlagInt",
	"DwarfAttEncodingEnum",
	"DwarfAttEncodingInt",
	"DwarfCCEnum",
	"DwarfCCInt",
	"DwarfLangEnum",
	"DwarfLangInt",
	"DwarfMacinfoEnum",
	"DwarfMacinfoInt",
	"DwarfOp",
	"DwarfTagEnum",
	"DwarfTagInt",
	"DwarfVirtualityEnum",
	"DwarfVirtualityInt",
	"EmissionKindEnum",
	"EmissionKindInt",
	"NameTableKindEnum",
	"NameTableKindInt",
	"AddrSpace",
	"Align",
	"AlignPair",
	"AlignStack",
	"AlignStackPair",
	"AllocSize",
	"Args",
	"Arg",
	"Atomic",
	"AtomicOrdering",
	"AttrPair",
	"AttrString",
	"Byval",
	"CallingConvEnum",
	"CallingConvInt",
	"Comdat",
	"Dereferenceable",
	"DereferenceableOrNull",
	"DLLStorageClass",
	"Ellipsis",
	"Exact",
	"ExceptionArg",
	"FastMathFlag",
	"FPred",
	"FuncAttr",
	"InBounds",
	"IPred",
	"Label",
	"Linkage",
	"ExternLinkage",
	"OperandBundle",
	"OverflowFlag",
	"Params",
	"Param",
	"ParamAttr",
	"Partition",
	"Preallocated",
	"Preemption",
	"ReturnAttr",
	"Section",
	"SyncScope",
	"ThreadLocal",
	"TLSModel",
	"TypeConst",
	"TypeValue",
	"UnnamedAddr",
	"UnwindToCaller",
	"Visibility",
	"Volatile",
}

func (t NodeType) String() string {
	if t >= 0 && int(t) < len(nodeTypeStr) {
		return nodeTypeStr[t]
	}
	return fmt.Sprintf("node(%d)", t)
}

var CallingConv = []NodeType{
	CallingConvEnum,
	CallingConvInt,
}

var ConcreteType = []NodeType{
	ArrayType,
	FloatType,
	IntType,
	LabelType,
	MMXType,
	NamedType,
	PackedStructType,
	PointerType,
	ScalableVectorType,
	StructType,
	TokenType,
	VectorType,
}

var Constant = []NodeType{
	AShrExpr,
	AddExpr,
	AddrSpaceCastExpr,
	AndExpr,
	ArrayConst,
	BitCastExpr,
	BlockAddressConst,
	BoolConst,
	CharArrayConst,
	ExtractElementExpr,
	ExtractValueExpr,
	FAddExpr,
	FCmpExpr,
	FDivExpr,
	FMulExpr,
	FNegExpr,
	FPExtExpr,
	FPToSIExpr,
	FPToUIExpr,
	FPTruncExpr,
	FRemExpr,
	FSubExpr,
	FloatConst,
	GetElementPtrExpr,
	GlobalIdent,
	ICmpExpr,
	InsertElementExpr,
	InsertValueExpr,
	IntConst,
	IntToPtrExpr,
	LShrExpr,
	MulExpr,
	NoneConst,
	NullConst,
	OrExpr,
	PtrToIntExpr,
	SDivExpr,
	SExtExpr,
	SIToFPExpr,
	SRemExpr,
	SelectExpr,
	ShlExpr,
	ShuffleVectorExpr,
	StructConst,
	SubExpr,
	TruncExpr,
	UDivExpr,
	UIToFPExpr,
	URemExpr,
	UndefConst,
	VectorConst,
	XorExpr,
	ZExtExpr,
	ZeroInitializerConst,
}

var ConstantExpr = []NodeType{
	AShrExpr,
	AddExpr,
	AddrSpaceCastExpr,
	AndExpr,
	BitCastExpr,
	ExtractElementExpr,
	ExtractValueExpr,
	FAddExpr,
	FCmpExpr,
	FDivExpr,
	FMulExpr,
	FNegExpr,
	FPExtExpr,
	FPToSIExpr,
	FPToUIExpr,
	FPTruncExpr,
	FRemExpr,
	FSubExpr,
	GetElementPtrExpr,
	ICmpExpr,
	InsertElementExpr,
	InsertValueExpr,
	IntToPtrExpr,
	LShrExpr,
	MulExpr,
	OrExpr,
	PtrToIntExpr,
	SDivExpr,
	SExtExpr,
	SIToFPExpr,
	SRemExpr,
	SelectExpr,
	ShlExpr,
	ShuffleVectorExpr,
	SubExpr,
	TruncExpr,
	UDivExpr,
	UIToFPExpr,
	URemExpr,
	XorExpr,
	ZExtExpr,
}

var DIBasicTypeField = []NodeType{
	AlignField,
	EncodingField,
	FlagsField,
	NameField,
	SizeField,
	TagField,
}

var DICommonBlockField = []NodeType{
	DeclarationField,
	FileField,
	LineField,
	NameField,
	ScopeField,
}

var DICompileUnitField = []NodeType{
	DebugInfoForProfilingField,
	DwoIdField,
	EmissionKindField,
	EnumsField,
	FileField,
	FlagsStringField,
	GlobalsField,
	ImportsField,
	IsOptimizedField,
	LanguageField,
	MacrosField,
	NameTableKindField,
	ProducerField,
	RangesBaseAddressField,
	RetainedTypesField,
	RuntimeVersionField,
	SDKField,
	SplitDebugFilenameField,
	SplitDebugInliningField,
	SysrootField,
}

var DICompositeTypeField = []NodeType{
	AlignField,
	BaseTypeField,
	DataLocationField,
	DiscriminatorField,
	ElementsField,
	FileField,
	FlagsField,
	IdentifierField,
	LineField,
	NameField,
	OffsetField,
	RuntimeLangField,
	ScopeField,
	SizeField,
	TagField,
	TemplateParamsField,
	VtableHolderField,
}

var DIDerivedTypeField = []NodeType{
	AlignField,
	BaseTypeField,
	DwarfAddressSpaceField,
	ExtraDataField,
	FileField,
	FlagsField,
	LineField,
	NameField,
	OffsetField,
	ScopeField,
	SizeField,
	TagField,
}

var DIEnumeratorField = []NodeType{
	IsUnsignedField,
	NameField,
	ValueIntField,
}

var DIExpressionField = []NodeType{
	DwarfAttEncodingEnum,
	DwarfOp,
	UintLit,
}

var DIFileField = []NodeType{
	ChecksumField,
	ChecksumkindField,
	DirectoryField,
	FilenameField,
	SourceField,
}

var DIFlag = []NodeType{
	DIFlagEnum,
	DIFlagInt,
}

var DIGlobalVariableExpressionField = []NodeType{
	ExprField,
	VarField,
}

var DIGlobalVariableField = []NodeType{
	AlignField,
	DeclarationField,
	FileField,
	IsDefinitionField,
	IsLocalField,
	LineField,
	LinkageNameField,
	NameField,
	ScopeField,
	TemplateParamsField,
	TypeField,
}

var DIImportedEntityField = []NodeType{
	EntityField,
	FileField,
	LineField,
	NameField,
	ScopeField,
	TagField,
}

var DILabelField = []NodeType{
	FileField,
	LineField,
	NameField,
	ScopeField,
}

var DILexicalBlockField = []NodeType{
	ColumnField,
	FileField,
	LineField,
	ScopeField,
}

var DILexicalBlockFileField = []NodeType{
	DiscriminatorIntField,
	FileField,
	ScopeField,
}

var DILocalVariableField = []NodeType{
	AlignField,
	ArgField,
	FileField,
	FlagsField,
	LineField,
	NameField,
	ScopeField,
	TypeField,
}

var DILocationField = []NodeType{
	ColumnField,
	InlinedAtField,
	IsImplicitCodeField,
	LineField,
	ScopeField,
}

var DIMacroField = []NodeType{
	LineField,
	NameField,
	TypeMacinfoField,
	ValueStringField,
}

var DIMacroFileField = []NodeType{
	FileField,
	LineField,
	NodesField,
	TypeMacinfoField,
}

var DIModuleField = []NodeType{
	APINotesField,
	ConfigMacrosField,
	FileField,
	IncludePathField,
	LineField,
	NameField,
	ScopeField,
}

var DINamespaceField = []NodeType{
	ExportSymbolsField,
	NameField,
	ScopeField,
}

var DIObjCPropertyField = []NodeType{
	AttributesField,
	FileField,
	GetterField,
	LineField,
	NameField,
	SetterField,
	TypeField,
}

var DISPFlag = []NodeType{
	DISPFlagEnum,
	DISPFlagInt,
}

var DISubprogramField = []NodeType{
	ContainingTypeField,
	DeclarationField,
	FileField,
	FlagsField,
	IsDefinitionField,
	IsLocalField,
	IsOptimizedField,
	LineField,
	LinkageNameField,
	NameField,
	RetainedNodesField,
	SPFlagsField,
	ScopeField,
	ScopeLineField,
	TemplateParamsField,
	ThisAdjustmentField,
	ThrownTypesField,
	TypeField,
	UnitField,
	VirtualIndexField,
	VirtualityField,
}

var DISubrangeField = []NodeType{
	CountField,
	LowerBoundField,
	StrideField,
	UpperBoundField,
}

var DISubroutineTypeField = []NodeType{
	CCField,
	FlagsField,
	TypesField,
}

var DITemplateTypeParameterField = []NodeType{
	DefaultedField,
	NameField,
	TypeField,
}

var DITemplateValueParameterField = []NodeType{
	DefaultedField,
	NameField,
	TagField,
	TypeField,
	ValueField,
}

var DwarfAttEncoding = []NodeType{
	DwarfAttEncodingEnum,
}

var DwarfAttEncodingOrUint = []NodeType{
	DwarfAttEncodingEnum,
	DwarfAttEncodingInt,
}

var DwarfCC = []NodeType{
	DwarfCCEnum,
	DwarfCCInt,
}

var DwarfLang = []NodeType{
	DwarfLangEnum,
	DwarfLangInt,
}

var DwarfMacinfo = []NodeType{
	DwarfMacinfoEnum,
	DwarfMacinfoInt,
}

var DwarfTag = []NodeType{
	DwarfTagEnum,
	DwarfTagInt,
}

var DwarfVirtuality = []NodeType{
	DwarfVirtualityEnum,
	DwarfVirtualityInt,
}

var EmissionKind = []NodeType{
	EmissionKindEnum,
	EmissionKindInt,
}

var ExceptionPad = []NodeType{
	LocalIdent,
	NoneConst,
}

var FirstClassType = []NodeType{
	ArrayType,
	FloatType,
	IntType,
	LabelType,
	MMXType,
	MetadataType,
	NamedType,
	PackedStructType,
	PointerType,
	ScalableVectorType,
	StructType,
	TokenType,
	VectorType,
}

var FuncAttribute = []NodeType{
	AlignPair,
	AlignStack,
	AlignStackPair,
	AllocSize,
	AttrGroupID,
	AttrPair,
	AttrString,
	FuncAttr,
	Preallocated,
}

var FuncHdrField = []NodeType{
	Align,
	AlignPair,
	AlignStack,
	AlignStackPair,
	AllocSize,
	AttrGroupID,
	AttrPair,
	AttrString,
	Comdat,
	FuncAttr,
	GCNode,
	Partition,
	Personality,
	Preallocated,
	Prefix,
	Prologue,
	Section,
}

var GenericDINodeField = []NodeType{
	HeaderField,
	OperandsField,
	TagField,
}

var GlobalField = []NodeType{
	Align,
	Comdat,
	Partition,
	Section,
}

var IndirectSymbol = []NodeType{
	AddrSpaceCastExpr,
	BitCastExpr,
	GetElementPtrExpr,
	IntToPtrExpr,
	TypeConst,
}

var Instruction = []NodeType{
	AShrInst,
	AddInst,
	AddrSpaceCastInst,
	AllocaInst,
	AndInst,
	AtomicRMWInst,
	BitCastInst,
	CallInst,
	CatchPadInst,
	CleanupPadInst,
	CmpXchgInst,
	ExtractElementInst,
	ExtractValueInst,
	FAddInst,
	FCmpInst,
	FDivInst,
	FMulInst,
	FNegInst,
	FPExtInst,
	FPToSIInst,
	FPToUIInst,
	FPTruncInst,
	FRemInst,
	FSubInst,
	FenceInst,
	FreezeInst,
	GetElementPtrInst,
	ICmpInst,
	InsertElementInst,
	InsertValueInst,
	IntToPtrInst,
	LShrInst,
	LandingPadInst,
	LoadInst,
	LocalDefInst,
	MulInst,
	OrInst,
	PhiInst,
	PtrToIntInst,
	SDivInst,
	SExtInst,
	SIToFPInst,
	SRemInst,
	SelectInst,
	ShlInst,
	ShuffleVectorInst,
	StoreInst,
	SubInst,
	TruncInst,
	UDivInst,
	UIToFPInst,
	URemInst,
	VAArgInst,
	XorInst,
	ZExtInst,
}

var MDField = []NodeType{
	DIBasicType,
	DICommonBlock,
	DICompileUnit,
	DICompositeType,
	DIDerivedType,
	DIEnumerator,
	DIExpression,
	DIFile,
	DIGlobalVariable,
	DIGlobalVariableExpression,
	DIImportedEntity,
	DILabel,
	DILexicalBlock,
	DILexicalBlockFile,
	DILocalVariable,
	DILocation,
	DIMacro,
	DIMacroFile,
	DIModule,
	DINamespace,
	DIObjCProperty,
	DISubprogram,
	DISubrange,
	DISubroutineType,
	DITemplateTypeParameter,
	DITemplateValueParameter,
	GenericDINode,
	MDString,
	MDTuple,
	MetadataID,
	NullLit,
	TypeValue,
}

var MDFieldOrInt = []NodeType{
	DIBasicType,
	DICommonBlock,
	DICompileUnit,
	DICompositeType,
	DIDerivedType,
	DIEnumerator,
	DIExpression,
	DIFile,
	DIGlobalVariable,
	DIGlobalVariableExpression,
	DIImportedEntity,
	DILabel,
	DILexicalBlock,
	DILexicalBlockFile,
	DILocalVariable,
	DILocation,
	DIMacro,
	DIMacroFile,
	DIModule,
	DINamespace,
	DIObjCProperty,
	DISubprogram,
	DISubrange,
	DISubroutineType,
	DITemplateTypeParameter,
	DITemplateValueParameter,
	GenericDINode,
	IntLit,
	MDString,
	MDTuple,
	MetadataID,
	NullLit,
	TypeValue,
}

var MDNode = []NodeType{
	DIBasicType,
	DICommonBlock,
	DICompileUnit,
	DICompositeType,
	DIDerivedType,
	DIEnumerator,
	DIExpression,
	DIFile,
	DIGlobalVariable,
	DIGlobalVariableExpression,
	DIImportedEntity,
	DILabel,
	DILexicalBlock,
	DILexicalBlockFile,
	DILocalVariable,
	DILocation,
	DIMacro,
	DIMacroFile,
	DIModule,
	DINamespace,
	DIObjCProperty,
	DISubprogram,
	DISubrange,
	DISubroutineType,
	DITemplateTypeParameter,
	DITemplateValueParameter,
	GenericDINode,
	MDTuple,
	MetadataID,
}

var Metadata = []NodeType{
	DIBasicType,
	DICommonBlock,
	DICompileUnit,
	DICompositeType,
	DIDerivedType,
	DIEnumerator,
	DIExpression,
	DIFile,
	DIGlobalVariable,
	DIGlobalVariableExpression,
	DIImportedEntity,
	DILabel,
	DILexicalBlock,
	DILexicalBlockFile,
	DILocalVariable,
	DILocation,
	DIMacro,
	DIMacroFile,
	DIModule,
	DINamespace,
	DIObjCProperty,
	DISubprogram,
	DISubrange,
	DISubroutineType,
	DITemplateTypeParameter,
	DITemplateValueParameter,
	GenericDINode,
	MDString,
	MDTuple,
	MetadataID,
	TypeValue,
}

var MetadataNode = []NodeType{
	DIExpression,
	MetadataID,
}

var NameTableKind = []NodeType{
	NameTableKindEnum,
	NameTableKindInt,
}

var ParamAttribute = []NodeType{
	Align,
	AttrPair,
	AttrString,
	Byval,
	Dereferenceable,
	DereferenceableOrNull,
	ParamAttr,
	Preallocated,
}

var ReturnAttribute = []NodeType{
	Dereferenceable,
	DereferenceableOrNull,
	ReturnAttr,
}

var SpecializedMDNode = []NodeType{
	DIBasicType,
	DICommonBlock,
	DICompileUnit,
	DICompositeType,
	DIDerivedType,
	DIEnumerator,
	DIExpression,
	DIFile,
	DIGlobalVariable,
	DIGlobalVariableExpression,
	DIImportedEntity,
	DILabel,
	DILexicalBlock,
	DILexicalBlockFile,
	DILocalVariable,
	DILocation,
	DIMacro,
	DIMacroFile,
	DIModule,
	DINamespace,
	DIObjCProperty,
	DISubprogram,
	DISubrange,
	DISubroutineType,
	DITemplateTypeParameter,
	DITemplateValueParameter,
	GenericDINode,
}

var TargetDef = []NodeType{
	SourceFilename,
	TargetDataLayout,
	TargetTriple,
}

var Terminator = []NodeType{
	BrTerm,
	CallBrTerm,
	CatchRetTerm,
	CatchSwitchTerm,
	CleanupRetTerm,
	CondBrTerm,
	IndirectBrTerm,
	InvokeTerm,
	LocalDefTerm,
	ResumeTerm,
	RetTerm,
	SwitchTerm,
	UnreachableTerm,
}

var TopLevelEntity = []NodeType{
	AttrGroupDef,
	ComdatDef,
	FuncDecl,
	FuncDef,
	GlobalDecl,
	IndirectSymbolDef,
	MetadataDef,
	ModuleAsm,
	NamedMetadataDef,
	TypeDef,
	UseListOrder,
	UseListOrderBB,
}

var Type = []NodeType{
	ArrayType,
	FloatType,
	FuncType,
	IntType,
	LabelType,
	MMXType,
	MetadataType,
	NamedType,
	PackedStructType,
	PointerType,
	ScalableVectorType,
	StructType,
	TokenType,
	VectorType,
	VoidType,
}

var UnwindTarget = []NodeType{
	Label,
	UnwindToCaller,
}

var Value = []NodeType{
	AShrExpr,
	AddExpr,
	AddrSpaceCastExpr,
	AndExpr,
	ArrayConst,
	BitCastExpr,
	BlockAddressConst,
	BoolConst,
	CharArrayConst,
	ExtractElementExpr,
	ExtractValueExpr,
	FAddExpr,
	FCmpExpr,
	FDivExpr,
	FMulExpr,
	FNegExpr,
	FPExtExpr,
	FPToSIExpr,
	FPToUIExpr,
	FPTruncExpr,
	FRemExpr,
	FSubExpr,
	FloatConst,
	GetElementPtrExpr,
	GlobalIdent,
	ICmpExpr,
	InlineAsm,
	InsertElementExpr,
	InsertValueExpr,
	IntConst,
	IntToPtrExpr,
	LShrExpr,
	LocalIdent,
	MulExpr,
	NoneConst,
	NullConst,
	OrExpr,
	PtrToIntExpr,
	SDivExpr,
	SExtExpr,
	SIToFPExpr,
	SRemExpr,
	SelectExpr,
	ShlExpr,
	ShuffleVectorExpr,
	StructConst,
	SubExpr,
	TruncExpr,
	UDivExpr,
	UIToFPExpr,
	URemExpr,
	UndefConst,
	VectorConst,
	XorExpr,
	ZExtExpr,
	ZeroInitializerConst,
}

var ValueInstruction = []NodeType{
	AShrInst,
	AddInst,
	AddrSpaceCastInst,
	AllocaInst,
	AndInst,
	AtomicRMWInst,
	BitCastInst,
	CallInst,
	CatchPadInst,
	CleanupPadInst,
	CmpXchgInst,
	ExtractElementInst,
	ExtractValueInst,
	FAddInst,
	FCmpInst,
	FDivInst,
	FMulInst,
	FNegInst,
	FPExtInst,
	FPToSIInst,
	FPToUIInst,
	FPTruncInst,
	FRemInst,
	FSubInst,
	FreezeInst,
	GetElementPtrInst,
	ICmpInst,
	InsertElementInst,
	InsertValueInst,
	IntToPtrInst,
	LShrInst,
	LandingPadInst,
	LoadInst,
	MulInst,
	OrInst,
	PhiInst,
	PtrToIntInst,
	SDivInst,
	SExtInst,
	SIToFPInst,
	SRemInst,
	SelectInst,
	ShlInst,
	ShuffleVectorInst,
	SubInst,
	TruncInst,
	UDivInst,
	UIToFPInst,
	URemInst,
	VAArgInst,
	XorInst,
	ZExtInst,
}

var ValueTerminator = []NodeType{
	CallBrTerm,
	CatchSwitchTerm,
	InvokeTerm,
}
