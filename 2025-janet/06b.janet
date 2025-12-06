# This file was written by Damien Dart, <damiendart@pobox.com>. This is
# free and unencumbered software released into the public domain. For
# more information, please refer to the accompanying "UNLICENCE" file.

(def input-parser (peg/compile ~{:main (* :numbers :ops)
                                 :ops (group (any (<- (* (set "*+") :s*))))
                                 :numbers (accumulate (any (<- (* :s* :d* :s*))))}))

(defn solve [input]
  (def [numbers ops] (peg/match input-parser input))

  (def values (do
                (def rows (string/split "\n" (string/trimr numbers "\n")))
                (var i 0)
                (map
                  (fn [j]
                    (def v (map |(string/slice $ i (+ i j)) rows))
                    (set i (+ i j))
                    v)
                  (map length ops))))

  (def rotated-values (map
                        (fn [v]
                          (def length (length (get v 0)))
                          (map (fn [i] (->> ;(map |(get $ i) v) string/from-bytes string/trim scan-number)) (range length)))
                        values))

  (do
    (var result 0)
    (eachp [k v] rotated-values
      (def numbers (filter number? v))
      (set result (+ result (case (string/trim (get ops k)) "+" (+ ;numbers) "*" (* ;numbers)))))
    result))

(defn main [&] (->> (string/trimr (file/read stdin :all) "\n") solve pp))
