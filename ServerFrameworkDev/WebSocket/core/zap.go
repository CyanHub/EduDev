package core

import (
	"ServerFramework/global"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapCore struct {
	level zapcore.Level
	zapcore.Core
}

func NewZapCore(level zapcore.Level) *ZapCore {
	entity := &ZapCore{level: level}
	syncer := entity.WriteSyncer()
	levelEnabler := zap.LevelEnablerFunc(func(l zapcore.Level) bool {
		return l == level
	})
	entity.Core = zapcore.NewCore(zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()), syncer, levelEnabler)
	return entity
}

func (z *ZapCore) WriteSyncer() zapcore.WriteSyncer {
	cutter := NewCutter(
		//CutterWithLayout("2006-01-02 15-04"),
		CutterWithLayout(global.CONFIG.Logger.Layout),
		CutterWithLevel(z.level),
		//CutterWithDirector("./sss"),
		CutterWithDirector(global.CONFIG.Logger.Director),
	)
	return zapcore.AddSync(cutter)
}
