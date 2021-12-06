<?php

	function isPangram(string $param): bool {
		
		$sentences = strtolower(trim($param));
		$letters = str_split("thequickbrownfoxjumpsoverthelazydog");

		foreach ($letters as $letter) {
			if (!strstr($sentences, $letter))
				return false;
		}

		return true;
	}