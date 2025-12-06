# This file was written by Damien Dart, <damiendart@pobox.com>. This is
# free and unencumbered software released into the public domain. For
# more information, please refer to the accompanying "UNLICENCE" file.

(def input-parser (peg/compile ~{:main (split :s+ :line)
                                 :line (any (+
                                              (number :d+)
                                              (<- (set "*+"))))}))

(defn solve [input]
  (def parsed (peg/match input-parser input))
  (var result 0)
  (let [ops (filter string? parsed)
        values (partition (length ops) (filter number? parsed))]
    (for i 0 (length ops) (set result (+ result (case (get ops i)
                                                  "*" (* ;(map |(get $ i) values))
                                                  "+" (+ ;(map |(get $ i) values)))))))
  result)

(defn main [&] (->> (file/read stdin :all) string/trim solve pp))
