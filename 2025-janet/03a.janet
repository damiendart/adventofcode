# This file was written by Damien Dart, <damiendart@pobox.com>. This is
# free and unencumbered software released into the public domain. For
# more information, please refer to the accompanying "UNLICENCE" file.

(def input-parser (peg/compile ~{:main (split :s (group :values))
                                 :values (* (some (number :d)))}))


(defn solve [input]
  (reduce
    (fn [carry bank]
      (def first-number (max ;(array/slice bank 0 -2)))
      (def second-number (max ;(array/slice bank (+ 1 (index-of first-number bank)))))
      (+ carry (* 10 first-number) second-number))
    0
    (peg/match input-parser input)))

(defn main [&] (->> (file/read stdin :all) string/trim solve pp))
