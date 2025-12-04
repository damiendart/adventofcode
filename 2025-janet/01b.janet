# This file was written by Damien Dart, <damiendart@pobox.com>. This is
# free and unencumbered software released into the public domain. For
# more information, please refer to the accompanying "UNLICENCE" file.

(def parser (peg/compile ~{:main (split :s (group :instruction))
                           :instruction (sequence (capture (set "LR")) (number :d+))}))

(defn solve [input]
  (var password 0)
  (reduce
    (fn [carry instruction]
      (var c carry)
      (for i 0 (get instruction 1)
        (if (= c 0) (set password (+ password 1)))
        (if (= (get instruction 0) "L")
          (set c (mod (-- c) 100))
          (set c (mod (++ c) 100))))
      c)
    50
    (peg/match parser input))
  password)

(defn main [&] (->> (file/read stdin :all) string/trim solve pp))
