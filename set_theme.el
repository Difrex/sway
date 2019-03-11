#!/usr/bin/emacs --script

;;; Code:

(defvar l (cdr (insert-file-contents-literally "~/.config/i3/config")))

(message l)

;;; set_theme.el ends here
