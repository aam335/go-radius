package radius

func init() {
	builtinOnce.Do(initDictionary)
	Builtin.MustRegister("Chargeable-User-Identity", 89, AttributeString)
}
