# This file was written by Damien Dart, <damiendart@pobox.com>. This is
# free and unencumbered software released into the public domain. For
# more information, please refer to the accompanying "UNLICENCE" file.

(def input-parser (peg/compile ~{:main (split "\n" (group (* (any (<- (set ".S^"))))))}))

(defn solve [input]
  (def parsed (peg/match input-parser input))
  (var lasers (array/new-filled (length (get parsed 0)) "."))

  (reduce
    (fn [carry line]
      (var splits 0)
      (eachp [idx column] line
        (case column
          "S" (put lasers idx "|")
          "^" (do
                (if (= (get lasers idx) "|") (do
                                               (++ splits)
                                               (put lasers idx ".")
                                               (if (not= (get lasers (- idx 1)) "|") (do
                                                                                       (put lasers (- idx 1) "|")))
                                               (if (not= (get lasers (+ idx 1)) "|") (do
                                                                                       (put lasers (+ idx 1) "|"))))))))
      (+ carry splits))
    0
    parsed))

(defn main [&] (->> (file/read stdin :all) string/trim solve pp))
