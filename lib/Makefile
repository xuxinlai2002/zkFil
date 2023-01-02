MAKE ?= $(shell which make)

ifndef V
  V := @
else
  V :=
  MAKE := V=1 $(MAKE)
endif

all: zk_setup zk_publish zk_core

zk_setup zk_publish zk_core:
	$(V) cd $@; $(MAKE)

.PHONY: zk_setup zk_publish zk_core
