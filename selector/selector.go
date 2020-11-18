// generated by Textmapper; DO NOT EDIT

package selector

import (
	"github.com/llir/ll"
)

type Selector func(nt ll.NodeType) bool

var (
	Any                             = func(t ll.NodeType) bool { return true }
	APINotesField                   = func(t ll.NodeType) bool { return t == ll.APINotesField }
	AShrExpr                        = func(t ll.NodeType) bool { return t == ll.AShrExpr }
	AShrInst                        = func(t ll.NodeType) bool { return t == ll.AShrInst }
	AddExpr                         = func(t ll.NodeType) bool { return t == ll.AddExpr }
	AddInst                         = func(t ll.NodeType) bool { return t == ll.AddInst }
	AddrSpace                       = func(t ll.NodeType) bool { return t == ll.AddrSpace }
	AddrSpaceCastExpr               = func(t ll.NodeType) bool { return t == ll.AddrSpaceCastExpr }
	AddrSpaceCastInst               = func(t ll.NodeType) bool { return t == ll.AddrSpaceCastInst }
	Align                           = func(t ll.NodeType) bool { return t == ll.Align }
	AlignField                      = func(t ll.NodeType) bool { return t == ll.AlignField }
	AlignPair                       = func(t ll.NodeType) bool { return t == ll.AlignPair }
	AlignStack                      = func(t ll.NodeType) bool { return t == ll.AlignStack }
	AlignStackPair                  = func(t ll.NodeType) bool { return t == ll.AlignStackPair }
	AlignStackTok                   = func(t ll.NodeType) bool { return t == ll.AlignStackTok }
	AllocSize                       = func(t ll.NodeType) bool { return t == ll.AllocSize }
	AllocaInst                      = func(t ll.NodeType) bool { return t == ll.AllocaInst }
	AndExpr                         = func(t ll.NodeType) bool { return t == ll.AndExpr }
	AndInst                         = func(t ll.NodeType) bool { return t == ll.AndInst }
	Arg                             = func(t ll.NodeType) bool { return t == ll.Arg }
	ArgField                        = func(t ll.NodeType) bool { return t == ll.ArgField }
	Args                            = func(t ll.NodeType) bool { return t == ll.Args }
	ArrayConst                      = func(t ll.NodeType) bool { return t == ll.ArrayConst }
	ArrayType                       = func(t ll.NodeType) bool { return t == ll.ArrayType }
	Atomic                          = func(t ll.NodeType) bool { return t == ll.Atomic }
	AtomicOp                        = func(t ll.NodeType) bool { return t == ll.AtomicOp }
	AtomicOrdering                  = func(t ll.NodeType) bool { return t == ll.AtomicOrdering }
	AtomicRMWInst                   = func(t ll.NodeType) bool { return t == ll.AtomicRMWInst }
	AttrGroupDef                    = func(t ll.NodeType) bool { return t == ll.AttrGroupDef }
	AttrGroupID                     = func(t ll.NodeType) bool { return t == ll.AttrGroupID }
	AttrPair                        = func(t ll.NodeType) bool { return t == ll.AttrPair }
	AttrString                      = func(t ll.NodeType) bool { return t == ll.AttrString }
	AttributesField                 = func(t ll.NodeType) bool { return t == ll.AttributesField }
	BaseTypeField                   = func(t ll.NodeType) bool { return t == ll.BaseTypeField }
	BasicBlock                      = func(t ll.NodeType) bool { return t == ll.BasicBlock }
	BitCastExpr                     = func(t ll.NodeType) bool { return t == ll.BitCastExpr }
	BitCastInst                     = func(t ll.NodeType) bool { return t == ll.BitCastInst }
	BlockAddressConst               = func(t ll.NodeType) bool { return t == ll.BlockAddressConst }
	BoolConst                       = func(t ll.NodeType) bool { return t == ll.BoolConst }
	BoolLit                         = func(t ll.NodeType) bool { return t == ll.BoolLit }
	BrTerm                          = func(t ll.NodeType) bool { return t == ll.BrTerm }
	Byval                           = func(t ll.NodeType) bool { return t == ll.Byval }
	CCField                         = func(t ll.NodeType) bool { return t == ll.CCField }
	CallBrTerm                      = func(t ll.NodeType) bool { return t == ll.CallBrTerm }
	CallInst                        = func(t ll.NodeType) bool { return t == ll.CallInst }
	CallingConvEnum                 = func(t ll.NodeType) bool { return t == ll.CallingConvEnum }
	CallingConvInt                  = func(t ll.NodeType) bool { return t == ll.CallingConvInt }
	Case                            = func(t ll.NodeType) bool { return t == ll.Case }
	CatchPadInst                    = func(t ll.NodeType) bool { return t == ll.CatchPadInst }
	CatchRetTerm                    = func(t ll.NodeType) bool { return t == ll.CatchRetTerm }
	CatchSwitchTerm                 = func(t ll.NodeType) bool { return t == ll.CatchSwitchTerm }
	CharArrayConst                  = func(t ll.NodeType) bool { return t == ll.CharArrayConst }
	ChecksumField                   = func(t ll.NodeType) bool { return t == ll.ChecksumField }
	ChecksumKind                    = func(t ll.NodeType) bool { return t == ll.ChecksumKind }
	ChecksumkindField               = func(t ll.NodeType) bool { return t == ll.ChecksumkindField }
	Clause                          = func(t ll.NodeType) bool { return t == ll.Clause }
	ClauseType                      = func(t ll.NodeType) bool { return t == ll.ClauseType }
	Cleanup                         = func(t ll.NodeType) bool { return t == ll.Cleanup }
	CleanupPadInst                  = func(t ll.NodeType) bool { return t == ll.CleanupPadInst }
	CleanupRetTerm                  = func(t ll.NodeType) bool { return t == ll.CleanupRetTerm }
	CmpXchgInst                     = func(t ll.NodeType) bool { return t == ll.CmpXchgInst }
	ColumnField                     = func(t ll.NodeType) bool { return t == ll.ColumnField }
	Comdat                          = func(t ll.NodeType) bool { return t == ll.Comdat }
	ComdatDef                       = func(t ll.NodeType) bool { return t == ll.ComdatDef }
	ComdatName                      = func(t ll.NodeType) bool { return t == ll.ComdatName }
	CondBrTerm                      = func(t ll.NodeType) bool { return t == ll.CondBrTerm }
	ConfigMacrosField               = func(t ll.NodeType) bool { return t == ll.ConfigMacrosField }
	ContainingTypeField             = func(t ll.NodeType) bool { return t == ll.ContainingTypeField }
	CountField                      = func(t ll.NodeType) bool { return t == ll.CountField }
	DIBasicType                     = func(t ll.NodeType) bool { return t == ll.DIBasicType }
	DICommonBlock                   = func(t ll.NodeType) bool { return t == ll.DICommonBlock }
	DICompileUnit                   = func(t ll.NodeType) bool { return t == ll.DICompileUnit }
	DICompositeType                 = func(t ll.NodeType) bool { return t == ll.DICompositeType }
	DIDerivedType                   = func(t ll.NodeType) bool { return t == ll.DIDerivedType }
	DIEnumerator                    = func(t ll.NodeType) bool { return t == ll.DIEnumerator }
	DIExpression                    = func(t ll.NodeType) bool { return t == ll.DIExpression }
	DIFile                          = func(t ll.NodeType) bool { return t == ll.DIFile }
	DIFlagEnum                      = func(t ll.NodeType) bool { return t == ll.DIFlagEnum }
	DIFlagInt                       = func(t ll.NodeType) bool { return t == ll.DIFlagInt }
	DIFlags                         = func(t ll.NodeType) bool { return t == ll.DIFlags }
	DIGlobalVariable                = func(t ll.NodeType) bool { return t == ll.DIGlobalVariable }
	DIGlobalVariableExpression      = func(t ll.NodeType) bool { return t == ll.DIGlobalVariableExpression }
	DIImportedEntity                = func(t ll.NodeType) bool { return t == ll.DIImportedEntity }
	DILabel                         = func(t ll.NodeType) bool { return t == ll.DILabel }
	DILexicalBlock                  = func(t ll.NodeType) bool { return t == ll.DILexicalBlock }
	DILexicalBlockFile              = func(t ll.NodeType) bool { return t == ll.DILexicalBlockFile }
	DILocalVariable                 = func(t ll.NodeType) bool { return t == ll.DILocalVariable }
	DILocation                      = func(t ll.NodeType) bool { return t == ll.DILocation }
	DIMacro                         = func(t ll.NodeType) bool { return t == ll.DIMacro }
	DIMacroFile                     = func(t ll.NodeType) bool { return t == ll.DIMacroFile }
	DIModule                        = func(t ll.NodeType) bool { return t == ll.DIModule }
	DINamespace                     = func(t ll.NodeType) bool { return t == ll.DINamespace }
	DIObjCProperty                  = func(t ll.NodeType) bool { return t == ll.DIObjCProperty }
	DISPFlagEnum                    = func(t ll.NodeType) bool { return t == ll.DISPFlagEnum }
	DISPFlagInt                     = func(t ll.NodeType) bool { return t == ll.DISPFlagInt }
	DISPFlags                       = func(t ll.NodeType) bool { return t == ll.DISPFlags }
	DISubprogram                    = func(t ll.NodeType) bool { return t == ll.DISubprogram }
	DISubrange                      = func(t ll.NodeType) bool { return t == ll.DISubrange }
	DISubroutineType                = func(t ll.NodeType) bool { return t == ll.DISubroutineType }
	DITemplateTypeParameter         = func(t ll.NodeType) bool { return t == ll.DITemplateTypeParameter }
	DITemplateValueParameter        = func(t ll.NodeType) bool { return t == ll.DITemplateValueParameter }
	DLLStorageClass                 = func(t ll.NodeType) bool { return t == ll.DLLStorageClass }
	DebugBaseAddressField           = func(t ll.NodeType) bool { return t == ll.DebugBaseAddressField }
	DebugInfoForProfilingField      = func(t ll.NodeType) bool { return t == ll.DebugInfoForProfilingField }
	DeclarationField                = func(t ll.NodeType) bool { return t == ll.DeclarationField }
	Dereferenceable                 = func(t ll.NodeType) bool { return t == ll.Dereferenceable }
	DereferenceableOrNull           = func(t ll.NodeType) bool { return t == ll.DereferenceableOrNull }
	DirectoryField                  = func(t ll.NodeType) bool { return t == ll.DirectoryField }
	DiscriminatorField              = func(t ll.NodeType) bool { return t == ll.DiscriminatorField }
	DiscriminatorIntField           = func(t ll.NodeType) bool { return t == ll.DiscriminatorIntField }
	Distinct                        = func(t ll.NodeType) bool { return t == ll.Distinct }
	DwarfAddressSpaceField          = func(t ll.NodeType) bool { return t == ll.DwarfAddressSpaceField }
	DwarfAttEncodingEnum            = func(t ll.NodeType) bool { return t == ll.DwarfAttEncodingEnum }
	DwarfAttEncodingInt             = func(t ll.NodeType) bool { return t == ll.DwarfAttEncodingInt }
	DwarfCCEnum                     = func(t ll.NodeType) bool { return t == ll.DwarfCCEnum }
	DwarfCCInt                      = func(t ll.NodeType) bool { return t == ll.DwarfCCInt }
	DwarfLangEnum                   = func(t ll.NodeType) bool { return t == ll.DwarfLangEnum }
	DwarfLangInt                    = func(t ll.NodeType) bool { return t == ll.DwarfLangInt }
	DwarfMacinfoEnum                = func(t ll.NodeType) bool { return t == ll.DwarfMacinfoEnum }
	DwarfMacinfoInt                 = func(t ll.NodeType) bool { return t == ll.DwarfMacinfoInt }
	DwarfOp                         = func(t ll.NodeType) bool { return t == ll.DwarfOp }
	DwarfTagEnum                    = func(t ll.NodeType) bool { return t == ll.DwarfTagEnum }
	DwarfTagInt                     = func(t ll.NodeType) bool { return t == ll.DwarfTagInt }
	DwarfVirtualityEnum             = func(t ll.NodeType) bool { return t == ll.DwarfVirtualityEnum }
	DwarfVirtualityInt              = func(t ll.NodeType) bool { return t == ll.DwarfVirtualityInt }
	DwoIdField                      = func(t ll.NodeType) bool { return t == ll.DwoIdField }
	ElementsField                   = func(t ll.NodeType) bool { return t == ll.ElementsField }
	Ellipsis                        = func(t ll.NodeType) bool { return t == ll.Ellipsis }
	EmissionKindEnum                = func(t ll.NodeType) bool { return t == ll.EmissionKindEnum }
	EmissionKindField               = func(t ll.NodeType) bool { return t == ll.EmissionKindField }
	EmissionKindInt                 = func(t ll.NodeType) bool { return t == ll.EmissionKindInt }
	EncodingField                   = func(t ll.NodeType) bool { return t == ll.EncodingField }
	EntityField                     = func(t ll.NodeType) bool { return t == ll.EntityField }
	EnumsField                      = func(t ll.NodeType) bool { return t == ll.EnumsField }
	Exact                           = func(t ll.NodeType) bool { return t == ll.Exact }
	ExceptionArg                    = func(t ll.NodeType) bool { return t == ll.ExceptionArg }
	ExportSymbolsField              = func(t ll.NodeType) bool { return t == ll.ExportSymbolsField }
	ExprField                       = func(t ll.NodeType) bool { return t == ll.ExprField }
	ExternLinkage                   = func(t ll.NodeType) bool { return t == ll.ExternLinkage }
	ExternallyInitialized           = func(t ll.NodeType) bool { return t == ll.ExternallyInitialized }
	ExtraDataField                  = func(t ll.NodeType) bool { return t == ll.ExtraDataField }
	ExtractElementExpr              = func(t ll.NodeType) bool { return t == ll.ExtractElementExpr }
	ExtractElementInst              = func(t ll.NodeType) bool { return t == ll.ExtractElementInst }
	ExtractValueExpr                = func(t ll.NodeType) bool { return t == ll.ExtractValueExpr }
	ExtractValueInst                = func(t ll.NodeType) bool { return t == ll.ExtractValueInst }
	FAddExpr                        = func(t ll.NodeType) bool { return t == ll.FAddExpr }
	FAddInst                        = func(t ll.NodeType) bool { return t == ll.FAddInst }
	FCmpExpr                        = func(t ll.NodeType) bool { return t == ll.FCmpExpr }
	FCmpInst                        = func(t ll.NodeType) bool { return t == ll.FCmpInst }
	FDivExpr                        = func(t ll.NodeType) bool { return t == ll.FDivExpr }
	FDivInst                        = func(t ll.NodeType) bool { return t == ll.FDivInst }
	FMulExpr                        = func(t ll.NodeType) bool { return t == ll.FMulExpr }
	FMulInst                        = func(t ll.NodeType) bool { return t == ll.FMulInst }
	FNegExpr                        = func(t ll.NodeType) bool { return t == ll.FNegExpr }
	FNegInst                        = func(t ll.NodeType) bool { return t == ll.FNegInst }
	FPExtExpr                       = func(t ll.NodeType) bool { return t == ll.FPExtExpr }
	FPExtInst                       = func(t ll.NodeType) bool { return t == ll.FPExtInst }
	FPToSIExpr                      = func(t ll.NodeType) bool { return t == ll.FPToSIExpr }
	FPToSIInst                      = func(t ll.NodeType) bool { return t == ll.FPToSIInst }
	FPToUIExpr                      = func(t ll.NodeType) bool { return t == ll.FPToUIExpr }
	FPToUIInst                      = func(t ll.NodeType) bool { return t == ll.FPToUIInst }
	FPTruncExpr                     = func(t ll.NodeType) bool { return t == ll.FPTruncExpr }
	FPTruncInst                     = func(t ll.NodeType) bool { return t == ll.FPTruncInst }
	FPred                           = func(t ll.NodeType) bool { return t == ll.FPred }
	FRemExpr                        = func(t ll.NodeType) bool { return t == ll.FRemExpr }
	FRemInst                        = func(t ll.NodeType) bool { return t == ll.FRemInst }
	FSubExpr                        = func(t ll.NodeType) bool { return t == ll.FSubExpr }
	FSubInst                        = func(t ll.NodeType) bool { return t == ll.FSubInst }
	FastMathFlag                    = func(t ll.NodeType) bool { return t == ll.FastMathFlag }
	FenceInst                       = func(t ll.NodeType) bool { return t == ll.FenceInst }
	FileField                       = func(t ll.NodeType) bool { return t == ll.FileField }
	FilenameField                   = func(t ll.NodeType) bool { return t == ll.FilenameField }
	FlagsField                      = func(t ll.NodeType) bool { return t == ll.FlagsField }
	FlagsStringField                = func(t ll.NodeType) bool { return t == ll.FlagsStringField }
	FloatConst                      = func(t ll.NodeType) bool { return t == ll.FloatConst }
	FloatKind                       = func(t ll.NodeType) bool { return t == ll.FloatKind }
	FloatLit                        = func(t ll.NodeType) bool { return t == ll.FloatLit }
	FloatType                       = func(t ll.NodeType) bool { return t == ll.FloatType }
	FreezeInst                      = func(t ll.NodeType) bool { return t == ll.FreezeInst }
	FuncAttr                        = func(t ll.NodeType) bool { return t == ll.FuncAttr }
	FuncBody                        = func(t ll.NodeType) bool { return t == ll.FuncBody }
	FuncDecl                        = func(t ll.NodeType) bool { return t == ll.FuncDecl }
	FuncDef                         = func(t ll.NodeType) bool { return t == ll.FuncDef }
	FuncHeader                      = func(t ll.NodeType) bool { return t == ll.FuncHeader }
	FuncType                        = func(t ll.NodeType) bool { return t == ll.FuncType }
	GCNode                          = func(t ll.NodeType) bool { return t == ll.GCNode }
	GEPIndex                        = func(t ll.NodeType) bool { return t == ll.GEPIndex }
	GenericDINode                   = func(t ll.NodeType) bool { return t == ll.GenericDINode }
	GetElementPtrExpr               = func(t ll.NodeType) bool { return t == ll.GetElementPtrExpr }
	GetElementPtrInst               = func(t ll.NodeType) bool { return t == ll.GetElementPtrInst }
	GetterField                     = func(t ll.NodeType) bool { return t == ll.GetterField }
	GlobalDecl                      = func(t ll.NodeType) bool { return t == ll.GlobalDecl }
	GlobalIdent                     = func(t ll.NodeType) bool { return t == ll.GlobalIdent }
	GlobalsField                    = func(t ll.NodeType) bool { return t == ll.GlobalsField }
	Handlers                        = func(t ll.NodeType) bool { return t == ll.Handlers }
	HeaderField                     = func(t ll.NodeType) bool { return t == ll.HeaderField }
	ICmpExpr                        = func(t ll.NodeType) bool { return t == ll.ICmpExpr }
	ICmpInst                        = func(t ll.NodeType) bool { return t == ll.ICmpInst }
	IPred                           = func(t ll.NodeType) bool { return t == ll.IPred }
	IdentifierField                 = func(t ll.NodeType) bool { return t == ll.IdentifierField }
	Immutable                       = func(t ll.NodeType) bool { return t == ll.Immutable }
	ImportsField                    = func(t ll.NodeType) bool { return t == ll.ImportsField }
	InAlloca                        = func(t ll.NodeType) bool { return t == ll.InAlloca }
	InBounds                        = func(t ll.NodeType) bool { return t == ll.InBounds }
	InRange                         = func(t ll.NodeType) bool { return t == ll.InRange }
	Inc                             = func(t ll.NodeType) bool { return t == ll.Inc }
	IncludePathField                = func(t ll.NodeType) bool { return t == ll.IncludePathField }
	IndirectBrTerm                  = func(t ll.NodeType) bool { return t == ll.IndirectBrTerm }
	IndirectSymbolDef               = func(t ll.NodeType) bool { return t == ll.IndirectSymbolDef }
	IndirectSymbolKind              = func(t ll.NodeType) bool { return t == ll.IndirectSymbolKind }
	InlineAsm                       = func(t ll.NodeType) bool { return t == ll.InlineAsm }
	InlinedAtField                  = func(t ll.NodeType) bool { return t == ll.InlinedAtField }
	InsertElementExpr               = func(t ll.NodeType) bool { return t == ll.InsertElementExpr }
	InsertElementInst               = func(t ll.NodeType) bool { return t == ll.InsertElementInst }
	InsertValueExpr                 = func(t ll.NodeType) bool { return t == ll.InsertValueExpr }
	InsertValueInst                 = func(t ll.NodeType) bool { return t == ll.InsertValueInst }
	IntConst                        = func(t ll.NodeType) bool { return t == ll.IntConst }
	IntLit                          = func(t ll.NodeType) bool { return t == ll.IntLit }
	IntToPtrExpr                    = func(t ll.NodeType) bool { return t == ll.IntToPtrExpr }
	IntToPtrInst                    = func(t ll.NodeType) bool { return t == ll.IntToPtrInst }
	IntType                         = func(t ll.NodeType) bool { return t == ll.IntType }
	IntelDialect                    = func(t ll.NodeType) bool { return t == ll.IntelDialect }
	InvokeTerm                      = func(t ll.NodeType) bool { return t == ll.InvokeTerm }
	IsDefinitionField               = func(t ll.NodeType) bool { return t == ll.IsDefinitionField }
	IsImplicitCodeField             = func(t ll.NodeType) bool { return t == ll.IsImplicitCodeField }
	IsLocalField                    = func(t ll.NodeType) bool { return t == ll.IsLocalField }
	IsOptimizedField                = func(t ll.NodeType) bool { return t == ll.IsOptimizedField }
	IsUnsignedField                 = func(t ll.NodeType) bool { return t == ll.IsUnsignedField }
	LShrExpr                        = func(t ll.NodeType) bool { return t == ll.LShrExpr }
	LShrInst                        = func(t ll.NodeType) bool { return t == ll.LShrInst }
	Label                           = func(t ll.NodeType) bool { return t == ll.Label }
	LabelIdent                      = func(t ll.NodeType) bool { return t == ll.LabelIdent }
	LabelType                       = func(t ll.NodeType) bool { return t == ll.LabelType }
	LandingPadInst                  = func(t ll.NodeType) bool { return t == ll.LandingPadInst }
	LanguageField                   = func(t ll.NodeType) bool { return t == ll.LanguageField }
	LineField                       = func(t ll.NodeType) bool { return t == ll.LineField }
	Linkage                         = func(t ll.NodeType) bool { return t == ll.Linkage }
	LinkageNameField                = func(t ll.NodeType) bool { return t == ll.LinkageNameField }
	LoadInst                        = func(t ll.NodeType) bool { return t == ll.LoadInst }
	LocalDefInst                    = func(t ll.NodeType) bool { return t == ll.LocalDefInst }
	LocalDefTerm                    = func(t ll.NodeType) bool { return t == ll.LocalDefTerm }
	LocalIdent                      = func(t ll.NodeType) bool { return t == ll.LocalIdent }
	LowerBoundField                 = func(t ll.NodeType) bool { return t == ll.LowerBoundField }
	MDString                        = func(t ll.NodeType) bool { return t == ll.MDString }
	MDTuple                         = func(t ll.NodeType) bool { return t == ll.MDTuple }
	MMXType                         = func(t ll.NodeType) bool { return t == ll.MMXType }
	MacrosField                     = func(t ll.NodeType) bool { return t == ll.MacrosField }
	MetadataAttachment              = func(t ll.NodeType) bool { return t == ll.MetadataAttachment }
	MetadataDef                     = func(t ll.NodeType) bool { return t == ll.MetadataDef }
	MetadataID                      = func(t ll.NodeType) bool { return t == ll.MetadataID }
	MetadataName                    = func(t ll.NodeType) bool { return t == ll.MetadataName }
	MetadataType                    = func(t ll.NodeType) bool { return t == ll.MetadataType }
	Module                          = func(t ll.NodeType) bool { return t == ll.Module }
	ModuleAsm                       = func(t ll.NodeType) bool { return t == ll.ModuleAsm }
	MulExpr                         = func(t ll.NodeType) bool { return t == ll.MulExpr }
	MulInst                         = func(t ll.NodeType) bool { return t == ll.MulInst }
	NameField                       = func(t ll.NodeType) bool { return t == ll.NameField }
	NameTableKindEnum               = func(t ll.NodeType) bool { return t == ll.NameTableKindEnum }
	NameTableKindField              = func(t ll.NodeType) bool { return t == ll.NameTableKindField }
	NameTableKindInt                = func(t ll.NodeType) bool { return t == ll.NameTableKindInt }
	NamedMetadataDef                = func(t ll.NodeType) bool { return t == ll.NamedMetadataDef }
	NamedType                       = func(t ll.NodeType) bool { return t == ll.NamedType }
	NodesField                      = func(t ll.NodeType) bool { return t == ll.NodesField }
	NoneConst                       = func(t ll.NodeType) bool { return t == ll.NoneConst }
	NullConst                       = func(t ll.NodeType) bool { return t == ll.NullConst }
	NullLit                         = func(t ll.NodeType) bool { return t == ll.NullLit }
	OffsetField                     = func(t ll.NodeType) bool { return t == ll.OffsetField }
	OpaqueType                      = func(t ll.NodeType) bool { return t == ll.OpaqueType }
	OperandBundle                   = func(t ll.NodeType) bool { return t == ll.OperandBundle }
	OperandsField                   = func(t ll.NodeType) bool { return t == ll.OperandsField }
	OrExpr                          = func(t ll.NodeType) bool { return t == ll.OrExpr }
	OrInst                          = func(t ll.NodeType) bool { return t == ll.OrInst }
	OverflowFlag                    = func(t ll.NodeType) bool { return t == ll.OverflowFlag }
	PackedStructType                = func(t ll.NodeType) bool { return t == ll.PackedStructType }
	Param                           = func(t ll.NodeType) bool { return t == ll.Param }
	ParamAttr                       = func(t ll.NodeType) bool { return t == ll.ParamAttr }
	Params                          = func(t ll.NodeType) bool { return t == ll.Params }
	Partition                       = func(t ll.NodeType) bool { return t == ll.Partition }
	Personality                     = func(t ll.NodeType) bool { return t == ll.Personality }
	PhiInst                         = func(t ll.NodeType) bool { return t == ll.PhiInst }
	PointerType                     = func(t ll.NodeType) bool { return t == ll.PointerType }
	Preemption                      = func(t ll.NodeType) bool { return t == ll.Preemption }
	Prefix                          = func(t ll.NodeType) bool { return t == ll.Prefix }
	ProducerField                   = func(t ll.NodeType) bool { return t == ll.ProducerField }
	Prologue                        = func(t ll.NodeType) bool { return t == ll.Prologue }
	PtrToIntExpr                    = func(t ll.NodeType) bool { return t == ll.PtrToIntExpr }
	PtrToIntInst                    = func(t ll.NodeType) bool { return t == ll.PtrToIntInst }
	ResumeTerm                      = func(t ll.NodeType) bool { return t == ll.ResumeTerm }
	RetTerm                         = func(t ll.NodeType) bool { return t == ll.RetTerm }
	RetainedNodesField              = func(t ll.NodeType) bool { return t == ll.RetainedNodesField }
	RetainedTypesField              = func(t ll.NodeType) bool { return t == ll.RetainedTypesField }
	ReturnAttr                      = func(t ll.NodeType) bool { return t == ll.ReturnAttr }
	RuntimeLangField                = func(t ll.NodeType) bool { return t == ll.RuntimeLangField }
	RuntimeVersionField             = func(t ll.NodeType) bool { return t == ll.RuntimeVersionField }
	SDivExpr                        = func(t ll.NodeType) bool { return t == ll.SDivExpr }
	SDivInst                        = func(t ll.NodeType) bool { return t == ll.SDivInst }
	SExtExpr                        = func(t ll.NodeType) bool { return t == ll.SExtExpr }
	SExtInst                        = func(t ll.NodeType) bool { return t == ll.SExtInst }
	SIToFPExpr                      = func(t ll.NodeType) bool { return t == ll.SIToFPExpr }
	SIToFPInst                      = func(t ll.NodeType) bool { return t == ll.SIToFPInst }
	SPFlagsField                    = func(t ll.NodeType) bool { return t == ll.SPFlagsField }
	SRemExpr                        = func(t ll.NodeType) bool { return t == ll.SRemExpr }
	SRemInst                        = func(t ll.NodeType) bool { return t == ll.SRemInst }
	ScalableVectorType              = func(t ll.NodeType) bool { return t == ll.ScalableVectorType }
	ScopeField                      = func(t ll.NodeType) bool { return t == ll.ScopeField }
	ScopeLineField                  = func(t ll.NodeType) bool { return t == ll.ScopeLineField }
	Section                         = func(t ll.NodeType) bool { return t == ll.Section }
	SelectExpr                      = func(t ll.NodeType) bool { return t == ll.SelectExpr }
	SelectInst                      = func(t ll.NodeType) bool { return t == ll.SelectInst }
	SelectionKind                   = func(t ll.NodeType) bool { return t == ll.SelectionKind }
	SetterField                     = func(t ll.NodeType) bool { return t == ll.SetterField }
	ShlExpr                         = func(t ll.NodeType) bool { return t == ll.ShlExpr }
	ShlInst                         = func(t ll.NodeType) bool { return t == ll.ShlInst }
	ShuffleVectorExpr               = func(t ll.NodeType) bool { return t == ll.ShuffleVectorExpr }
	ShuffleVectorInst               = func(t ll.NodeType) bool { return t == ll.ShuffleVectorInst }
	SideEffect                      = func(t ll.NodeType) bool { return t == ll.SideEffect }
	SizeField                       = func(t ll.NodeType) bool { return t == ll.SizeField }
	SourceField                     = func(t ll.NodeType) bool { return t == ll.SourceField }
	SourceFilename                  = func(t ll.NodeType) bool { return t == ll.SourceFilename }
	SplitDebugFilenameField         = func(t ll.NodeType) bool { return t == ll.SplitDebugFilenameField }
	SplitDebugInliningField         = func(t ll.NodeType) bool { return t == ll.SplitDebugInliningField }
	StoreInst                       = func(t ll.NodeType) bool { return t == ll.StoreInst }
	StrideField                     = func(t ll.NodeType) bool { return t == ll.StrideField }
	StringLit                       = func(t ll.NodeType) bool { return t == ll.StringLit }
	StructConst                     = func(t ll.NodeType) bool { return t == ll.StructConst }
	StructType                      = func(t ll.NodeType) bool { return t == ll.StructType }
	SubExpr                         = func(t ll.NodeType) bool { return t == ll.SubExpr }
	SubInst                         = func(t ll.NodeType) bool { return t == ll.SubInst }
	SwiftError                      = func(t ll.NodeType) bool { return t == ll.SwiftError }
	SwitchTerm                      = func(t ll.NodeType) bool { return t == ll.SwitchTerm }
	SyncScope                       = func(t ll.NodeType) bool { return t == ll.SyncScope }
	TLSModel                        = func(t ll.NodeType) bool { return t == ll.TLSModel }
	TagField                        = func(t ll.NodeType) bool { return t == ll.TagField }
	Tail                            = func(t ll.NodeType) bool { return t == ll.Tail }
	TargetDataLayout                = func(t ll.NodeType) bool { return t == ll.TargetDataLayout }
	TargetTriple                    = func(t ll.NodeType) bool { return t == ll.TargetTriple }
	TemplateParamsField             = func(t ll.NodeType) bool { return t == ll.TemplateParamsField }
	ThisAdjustmentField             = func(t ll.NodeType) bool { return t == ll.ThisAdjustmentField }
	ThreadLocal                     = func(t ll.NodeType) bool { return t == ll.ThreadLocal }
	ThrownTypesField                = func(t ll.NodeType) bool { return t == ll.ThrownTypesField }
	TokenType                       = func(t ll.NodeType) bool { return t == ll.TokenType }
	TruncExpr                       = func(t ll.NodeType) bool { return t == ll.TruncExpr }
	TruncInst                       = func(t ll.NodeType) bool { return t == ll.TruncInst }
	TypeConst                       = func(t ll.NodeType) bool { return t == ll.TypeConst }
	TypeDef                         = func(t ll.NodeType) bool { return t == ll.TypeDef }
	TypeField                       = func(t ll.NodeType) bool { return t == ll.TypeField }
	TypeMacinfoField                = func(t ll.NodeType) bool { return t == ll.TypeMacinfoField }
	TypeValue                       = func(t ll.NodeType) bool { return t == ll.TypeValue }
	TypesField                      = func(t ll.NodeType) bool { return t == ll.TypesField }
	UDivExpr                        = func(t ll.NodeType) bool { return t == ll.UDivExpr }
	UDivInst                        = func(t ll.NodeType) bool { return t == ll.UDivInst }
	UIToFPExpr                      = func(t ll.NodeType) bool { return t == ll.UIToFPExpr }
	UIToFPInst                      = func(t ll.NodeType) bool { return t == ll.UIToFPInst }
	URemExpr                        = func(t ll.NodeType) bool { return t == ll.URemExpr }
	URemInst                        = func(t ll.NodeType) bool { return t == ll.URemInst }
	UintLit                         = func(t ll.NodeType) bool { return t == ll.UintLit }
	UndefConst                      = func(t ll.NodeType) bool { return t == ll.UndefConst }
	UnitField                       = func(t ll.NodeType) bool { return t == ll.UnitField }
	UnnamedAddr                     = func(t ll.NodeType) bool { return t == ll.UnnamedAddr }
	UnreachableTerm                 = func(t ll.NodeType) bool { return t == ll.UnreachableTerm }
	UnwindToCaller                  = func(t ll.NodeType) bool { return t == ll.UnwindToCaller }
	UpperBoundField                 = func(t ll.NodeType) bool { return t == ll.UpperBoundField }
	UseListOrder                    = func(t ll.NodeType) bool { return t == ll.UseListOrder }
	UseListOrderBB                  = func(t ll.NodeType) bool { return t == ll.UseListOrderBB }
	VAArgInst                       = func(t ll.NodeType) bool { return t == ll.VAArgInst }
	ValueField                      = func(t ll.NodeType) bool { return t == ll.ValueField }
	ValueIntField                   = func(t ll.NodeType) bool { return t == ll.ValueIntField }
	ValueStringField                = func(t ll.NodeType) bool { return t == ll.ValueStringField }
	VarField                        = func(t ll.NodeType) bool { return t == ll.VarField }
	VectorConst                     = func(t ll.NodeType) bool { return t == ll.VectorConst }
	VectorType                      = func(t ll.NodeType) bool { return t == ll.VectorType }
	VirtualIndexField               = func(t ll.NodeType) bool { return t == ll.VirtualIndexField }
	VirtualityField                 = func(t ll.NodeType) bool { return t == ll.VirtualityField }
	Visibility                      = func(t ll.NodeType) bool { return t == ll.Visibility }
	VoidType                        = func(t ll.NodeType) bool { return t == ll.VoidType }
	Volatile                        = func(t ll.NodeType) bool { return t == ll.Volatile }
	VtableHolderField               = func(t ll.NodeType) bool { return t == ll.VtableHolderField }
	Weak                            = func(t ll.NodeType) bool { return t == ll.Weak }
	XorExpr                         = func(t ll.NodeType) bool { return t == ll.XorExpr }
	XorInst                         = func(t ll.NodeType) bool { return t == ll.XorInst }
	ZExtExpr                        = func(t ll.NodeType) bool { return t == ll.ZExtExpr }
	ZExtInst                        = func(t ll.NodeType) bool { return t == ll.ZExtInst }
	ZeroInitializerConst            = func(t ll.NodeType) bool { return t == ll.ZeroInitializerConst }
	CallingConv                     = OneOf(ll.CallingConv...)
	ConcreteType                    = OneOf(ll.ConcreteType...)
	Constant                        = OneOf(ll.Constant...)
	ConstantExpr                    = OneOf(ll.ConstantExpr...)
	DIBasicTypeField                = OneOf(ll.DIBasicTypeField...)
	DICommonBlockField              = OneOf(ll.DICommonBlockField...)
	DICompileUnitField              = OneOf(ll.DICompileUnitField...)
	DICompositeTypeField            = OneOf(ll.DICompositeTypeField...)
	DIDerivedTypeField              = OneOf(ll.DIDerivedTypeField...)
	DIEnumeratorField               = OneOf(ll.DIEnumeratorField...)
	DIExpressionField               = OneOf(ll.DIExpressionField...)
	DIFileField                     = OneOf(ll.DIFileField...)
	DIFlag                          = OneOf(ll.DIFlag...)
	DIGlobalVariableExpressionField = OneOf(ll.DIGlobalVariableExpressionField...)
	DIGlobalVariableField           = OneOf(ll.DIGlobalVariableField...)
	DIImportedEntityField           = OneOf(ll.DIImportedEntityField...)
	DILabelField                    = OneOf(ll.DILabelField...)
	DILexicalBlockField             = OneOf(ll.DILexicalBlockField...)
	DILexicalBlockFileField         = OneOf(ll.DILexicalBlockFileField...)
	DILocalVariableField            = OneOf(ll.DILocalVariableField...)
	DILocationField                 = OneOf(ll.DILocationField...)
	DIMacroField                    = OneOf(ll.DIMacroField...)
	DIMacroFileField                = OneOf(ll.DIMacroFileField...)
	DIModuleField                   = OneOf(ll.DIModuleField...)
	DINamespaceField                = OneOf(ll.DINamespaceField...)
	DIObjCPropertyField             = OneOf(ll.DIObjCPropertyField...)
	DISPFlag                        = OneOf(ll.DISPFlag...)
	DISubprogramField               = OneOf(ll.DISubprogramField...)
	DISubrangeField                 = OneOf(ll.DISubrangeField...)
	DISubroutineTypeField           = OneOf(ll.DISubroutineTypeField...)
	DITemplateTypeParameterField    = OneOf(ll.DITemplateTypeParameterField...)
	DITemplateValueParameterField   = OneOf(ll.DITemplateValueParameterField...)
	DwarfAttEncoding                = OneOf(ll.DwarfAttEncoding...)
	DwarfAttEncodingOrUint          = OneOf(ll.DwarfAttEncodingOrUint...)
	DwarfCC                         = OneOf(ll.DwarfCC...)
	DwarfLang                       = OneOf(ll.DwarfLang...)
	DwarfMacinfo                    = OneOf(ll.DwarfMacinfo...)
	DwarfTag                        = OneOf(ll.DwarfTag...)
	DwarfVirtuality                 = OneOf(ll.DwarfVirtuality...)
	EmissionKind                    = OneOf(ll.EmissionKind...)
	ExceptionPad                    = OneOf(ll.ExceptionPad...)
	FirstClassType                  = OneOf(ll.FirstClassType...)
	FuncAttribute                   = OneOf(ll.FuncAttribute...)
	FuncHdrField                    = OneOf(ll.FuncHdrField...)
	GenericDINodeField              = OneOf(ll.GenericDINodeField...)
	GlobalField                     = OneOf(ll.GlobalField...)
	IndirectSymbol                  = OneOf(ll.IndirectSymbol...)
	Instruction                     = OneOf(ll.Instruction...)
	MDField                         = OneOf(ll.MDField...)
	MDFieldOrInt                    = OneOf(ll.MDFieldOrInt...)
	MDNode                          = OneOf(ll.MDNode...)
	Metadata                        = OneOf(ll.Metadata...)
	MetadataNode                    = OneOf(ll.MetadataNode...)
	NameTableKind                   = OneOf(ll.NameTableKind...)
	ParamAttribute                  = OneOf(ll.ParamAttribute...)
	ReturnAttribute                 = OneOf(ll.ReturnAttribute...)
	SpecializedMDNode               = OneOf(ll.SpecializedMDNode...)
	TargetDef                       = OneOf(ll.TargetDef...)
	Terminator                      = OneOf(ll.Terminator...)
	TopLevelEntity                  = OneOf(ll.TopLevelEntity...)
	Type                            = OneOf(ll.Type...)
	UnwindTarget                    = OneOf(ll.UnwindTarget...)
	Value                           = OneOf(ll.Value...)
	ValueInstruction                = OneOf(ll.ValueInstruction...)
	ValueTerminator                 = OneOf(ll.ValueTerminator...)
)

func OneOf(types ...ll.NodeType) Selector {
	if len(types) == 0 {
		return func(ll.NodeType) bool { return false }
	}
	const bits = 32
	max := 1
	for _, t := range types {
		if int(t) > max {
			max = int(t)
		}
	}
	size := (max + bits) / bits
	bitarr := make([]uint32, size)
	for _, t := range types {
		bitarr[uint(t)/bits] |= 1 << (uint(t) % bits)
	}
	return func(t ll.NodeType) bool {
		i := uint(t) / bits
		return int(i) < len(bitarr) && bitarr[i]&(1<<(uint(t)%bits)) != 0
	}
}
